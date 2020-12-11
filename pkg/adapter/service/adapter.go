package adapter

import (
	//"fmt"
	//"os"
	//"strings"

	"github.com/BrobridgeOrg/gravity-adapter-rest/pkg/app"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Adapter struct {
	app        app.App
	sm         *SourceManager
	clientName string
}

func NewAdapter(a app.App) *Adapter {
	adapter := &Adapter{
		app: a,
	}

	adapter.sm = NewSourceManager(adapter)

	return adapter
}

func (adapter *Adapter) Init() error {

	err := adapter.sm.Initialize()
	if err != nil {
		log.Error(err)
		return nil
	}

	return nil
}
