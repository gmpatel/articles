package service

import (
	"net/http"
)

// Service models the Sports Import Service
type Service struct {
	httpServer *http.Server
}

// Settings to load server settings from cli params
type Settings struct {
	ListenPort int
	ConnString string
	QryTimeout int
	Workers    int
}
