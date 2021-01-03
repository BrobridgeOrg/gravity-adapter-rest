package adapter

import (
	"time"
	//	"fmt"
	"net/http"
	"sync"
	"unsafe"

	dsa "github.com/BrobridgeOrg/gravity-api/service/dsa"
	parallel_chunked_flow "github.com/cfsghost/parallel-chunked-flow"
	"github.com/gin-gonic/gin"

	//	validation "github.com/go-ozzo/ozzo-validation"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

var counter uint64

type Packet struct {
	EventName string      `json:"event"`
	Payload   interface{} `json:"payload"`
}

type Source struct {
	adapter   *Adapter
	incoming  chan []byte
	name      string
	uri       string
	ginEngine *gin.Engine
	parser    *parallel_chunked_flow.ParallelChunkedFlow
}

var requestPool = sync.Pool{
	New: func() interface{} {
		return &dsa.PublishRequest{}
	},
}

func StrToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func NewSource(adapter *Adapter, name string, sourceInfo *SourceInfo) *Source {

	// required uri
	if len(sourceInfo.URI) == 0 {
		log.WithFields(log.Fields{
			"source": name,
		}).Error("Required uri")

		return nil
	}

	info := sourceInfo

	// Initialize parapllel chunked flow
	pcfOpts := parallel_chunked_flow.Options{
		BufferSize: 204800,
		ChunkSize:  512,
		ChunkCount: 512,
		Handler: func(data interface{}, output chan interface{}) {
			/*
				id := atomic.AddUint64((*uint64)(&counter), 1)
				if id%1000 == 0 {
					log.Info(id)
				}
			*/

			eventName := jsoniter.Get(data.([]byte), "event").ToString()
			payload := jsoniter.Get(data.([]byte), "payload").ToString()

			// Preparing request
			request := requestPool.Get().(*dsa.PublishRequest)
			request.EventName = eventName
			request.Payload = StrToBytes(payload)

			output <- request
		},
	}

	return &Source{
		adapter:  adapter,
		incoming: make(chan []byte, 204800),
		name:     name,
		uri:      info.URI,
		parser:   parallel_chunked_flow.NewParallelChunkedFlow(&pcfOpts),
	}
}

func (source *Source) InitSubscription() error {

	log.WithFields(log.Fields{
		"source": source.name,
	}).Info("Initializing Restful API.")

	source.ginEngine = source.adapter.app.GetHTTPServer().GetEngine()

	// API func
	source.ginEngine.POST(source.uri, func(c *gin.Context) {
		packet, err := c.GetRawData()
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		/*
			var body Packet
			err := c.BindJSON(&body)
			if err != nil {
				log.Error(err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			// Validate fields
			err = validation.ValidateStruct(&body,
				validation.Field(&body.EventName, validation.Required),
				validation.Field(&body.Payload, validation.Required),
			)
			if err != nil {
				log.Error(err)
				c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			//convert to []byte
			packet, err := json.Marshal(body)
			if err != nil {
				log.Error(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		*/

		// send message
		source.incoming <- packet

		// res
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})

	})

	return nil
}

func (source *Source) Init() error {

	log.WithFields(log.Fields{
		"source": source.name,
		"uri":    source.uri,
	}).Info("Initializing source connector")

	go source.eventReceiver()
	go source.requestHandler()

	return source.InitSubscription()
}

func (source *Source) eventReceiver() {

	log.WithFields(log.Fields{
		"source": source.name,
	}).Info("Initializing workers ...")

	for {
		select {
		case msg := <-source.incoming:
			source.parser.Push(msg)
		}
	}
}

func (source *Source) requestHandler() {

	for {
		select {
		case req := <-source.parser.Output():
			source.HandleRequest(req.(*dsa.PublishRequest))
			requestPool.Put(req)
		}
	}
}

func (source *Source) HandleRequest(request *dsa.PublishRequest) {

	for {
		connector := source.adapter.app.GetAdapterConnector()
		err := connector.Publish(request.EventName, request.Payload, nil)
		if err != nil {
			log.Error(err)
			time.Sleep(time.Second)
			continue
		}

		break
	}
}
