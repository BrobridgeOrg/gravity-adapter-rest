module github.com/BrobridgeOrg/gravity-adapter-rest

go 1.15

require (
	github.com/BrobridgeOrg/gravity-api v0.2.10
	github.com/BrobridgeOrg/gravity-sdk v0.0.2
	github.com/cfsghost/parallel-chunked-flow v0.0.2
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/json-iterator/go v1.1.10
	github.com/sirupsen/logrus v1.7.0
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/viper v1.7.1
	google.golang.org/grpc v1.32.0
//	google.golang.org/grpc v1.31.1
)

//replace github.com/BrobridgeOrg/gravity-api => ../gravity-api

//replace github.com/cfsghost/grpc-connection-pool => /Users/fred/works/opensource/grpc-connection-pool
