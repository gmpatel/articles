package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gmpatel/articles/cmd/article-api/app"
	"github.com/gmpatel/articles/controller"
	"github.com/gmpatel/articles/endpoint"
	"github.com/gmpatel/articles/repository"
	"github.com/gmpatel/articles/service"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	serviceSettings := service.Settings{}
	application := app.NewApp(&serviceSettings)

	application.Action = func(c *cli.Context) error {
		// Set log level
		log := setLooger()
		fmt.Println(serviceSettings.ConnString)
		// Construct our db repository
		repo, err := repository.NewRepository(serviceSettings.ConnString, serviceSettings.Workers, (serviceSettings.QryTimeout * 1000))
		if err != nil {
			return errors.Wrap(err, "Could not init the db repository")
		}

		// Setting up endpoints HTTP server
		httpServer := endpoint.NewEndpointServer(serviceSettings.ListenPort, repo)

		// Setting up service object with all necessary service components
		svc := service.NewService(httpServer)

		// Starting service
		svc.Start()

		// Waiting for the stop signal
		signals := make(chan os.Signal, 1)
		signal.Notify(signals,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)
		<-signals
		log.Info("Shutdown requested (^C OS Signal or similar)...")

		// Stopping service
		svc.Stop()

		// Return
		return nil
	}

	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setLooger() *logrus.Logger {
	// Prepare logger with the settings
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = os.Stdout

	level, errLevel := logrus.ParseLevel("DEBUG")
	if errLevel != nil {
		fmt.Println("Error logrus.ParseLevel('DEBUG')")
	}
	log.SetLevel(level)

	// Set logger on all packages we are going to use in this app
	app.SetLogger(log)
	repository.SetLogger(log)
	service.SetLogger(log)
	controller.SetLogger(log)
	endpoint.SetLogger(log)

	return log
}
