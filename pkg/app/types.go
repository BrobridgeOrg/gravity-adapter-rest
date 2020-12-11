package app

import (
	http_server "github.com/BrobridgeOrg/gravity-adapter-rest/pkg/http_server"
	mux_manager "github.com/BrobridgeOrg/gravity-adapter-rest/pkg/mux_manager"
	grpc_connection_pool "github.com/cfsghost/grpc-connection-pool"
)

type App interface {
	GetGRPCPool() *grpc_connection_pool.GRPCPool
	GetMuxManager() mux_manager.Manager
	GetHTTPServer() http_server.Server
}
