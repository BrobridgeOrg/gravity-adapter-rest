package instance

import (
	adapter_service "github.com/BrobridgeOrg/gravity-adapter-rest/pkg/adapter/service"
	http_server "github.com/BrobridgeOrg/gravity-adapter-rest/pkg/http_server/server"
	mux_manager "github.com/BrobridgeOrg/gravity-adapter-rest/pkg/mux_manager/manager"
	gravity_adapter "github.com/BrobridgeOrg/gravity-sdk/adapter"
	log "github.com/sirupsen/logrus"
)

type AppInstance struct {
	done             chan bool
	adapter          *adapter_service.Adapter
	adapterConnector *gravity_adapter.AdapterConnector
	httpServer       *http_server.Server
	muxManager       *mux_manager.MuxManager
}

func NewAppInstance() *AppInstance {

	a := &AppInstance{
		done: make(chan bool),
	}

	a.adapter = adapter_service.NewAdapter(a)

	return a
}

func (a *AppInstance) Init() error {

	a.muxManager = mux_manager.NewMuxManager(a)
	a.initMuxManager()
	a.httpServer = http_server.NewServer(a)

	// Initializing adapter connector
	err := a.initAdapterConnector()
	if err != nil {
		return err
	}

	// Initializing HTTP server
	err = a.initHTTPServer()
	if err != nil {
		return err
	}

	err = a.adapter.Init()
	if err != nil {
		return err
	}

	return nil
}

func (a *AppInstance) Uninit() {
}

func (a *AppInstance) Run() error {
	// HTTP
	go func() {
		err := a.runHTTPServer()
		if err != nil {
			log.Error(err)
		}
	}()

	err := a.runMuxManager()
	if err != nil {
		return err
	}

	<-a.done

	return nil
}
