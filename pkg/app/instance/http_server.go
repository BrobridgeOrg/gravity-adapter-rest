package instance

import (
	"errors"
	"fmt"

	http_server "github.com/BrobridgeOrg/gravity-adapter-rest/pkg/http_server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (a *AppInstance) initHTTPServer() error {

	// expose port
	port := viper.GetInt("service.port")
	host := fmt.Sprintf(":%d", port)

	log.Println(port)
	log.Println(host)

	if port == 0 {
		return errors.New("Required service port")
	}

	// Initializing HTTP server
	a.httpServer.Init(host)

	return nil
}

func (a *AppInstance) runHTTPServer() error {
	err := a.httpServer.Serve()
	if err != nil {
		log.Error(err)
		return err
	}

	return err
}

func (a *AppInstance) GetHTTPServer() http_server.Server {
	return http_server.Server(a.httpServer)
}
