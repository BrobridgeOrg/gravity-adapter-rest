package app

import (
	http_server "github.com/BrobridgeOrg/gravity-adapter-rest/pkg/http_server"
	mux_manager "github.com/BrobridgeOrg/gravity-adapter-rest/pkg/mux_manager"
	gravity_adapter "github.com/BrobridgeOrg/gravity-sdk/adapter"
)

type App interface {
	GetAdapterConnector() *gravity_adapter.AdapterConnector
	GetMuxManager() mux_manager.Manager
	GetHTTPServer() http_server.Server
}
