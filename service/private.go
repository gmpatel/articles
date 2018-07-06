package service

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func init() {
	SetLogger(logrus.StandardLogger())
}

func (svc *Service) startEndpoints() {
	log.Info("Attempting to start endpoints http listener...")
	err := svc.httpServer.ListenAndServe()
	if err != nil {
		log.WithFields(logrus.Fields{"error": err.Error()}).Error("Endpoints http listener failed to start...")
	}
}

func (svc *Service) stopEndpoints() {
	if svc.httpServer != nil {
		log.Infof("Attemptin to stop endpoints http listener...")
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
		defer cancel()

		errShutdown := svc.httpServer.Shutdown(ctx)
		if errShutdown != nil {
			log.WithError(errShutdown).Error("Error stopping endpoints http listener...")
		} else {
			log.Info("Endpoints http listener stopped...")
		}
	}
}
