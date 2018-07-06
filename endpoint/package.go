package endpoint

import (
	"fmt"
	"net/http"

	"github.com/articles"
	"github.com/sirupsen/logrus"
)

//NewEndpointServer returns the http server with endpoints ready to talk
func NewEndpointServer(listenPort int, repository articles.Repository) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%v", listenPort),
		Handler: setupRouter(repository),
	}
}

// SetLogger sets the logger for this package
func SetLogger(logger *logrus.Logger) {
	log = logger
}
