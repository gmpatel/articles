package service

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// NewService creates a new instance of the sports import service, ready for usage.
func NewService(httpServer *http.Server) *Service {
	return &Service{httpServer}
}

// Start will start the service, including its pollers and http server. This may wait if the repository cannot connect to the db
func (svc *Service) Start() {
	go svc.startEndpoints()
}

// Stop stops the service
func (svc *Service) Stop() {
	svc.stopEndpoints()
}

// SetLogger sets the logger for this package
func SetLogger(logger *logrus.Logger) {
	log = logger
}
