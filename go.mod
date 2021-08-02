module github.com/BrobridgeOrg/gravity-adapter-rest

go 1.15

require (
	github.com/BrobridgeOrg/gravity-sdk v0.0.31
	github.com/cfsghost/parallel-chunked-flow v0.0.6
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/json-iterator/go v1.1.10
	github.com/sirupsen/logrus v1.8.1
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/viper v1.7.1
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
)

//replace github.com/BrobridgeOrg/gravity-api => ../gravity-api

//replace github.com/cfsghost/grpc-connection-pool => /Users/fred/works/opensource/grpc-connection-pool
