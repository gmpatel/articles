package app

import (
	"github.com/gmpatel/articles/service"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

//DefaultSettings will return default app settings to be set
func DefaultSettings() *Settings {
	return &Settings{
		Name:        "article-api",
		Version:     "0.0.1",
		Description: "Article Service - serves the api related to the article app",
		Usage:       "Call the HTTP/HTTPS api endpoints as per the documentation provided",
	}
}

//DefaultServiceSettings will return default service settings before reading new values from the app cli
func DefaultServiceSettings() *service.Settings {
	return &service.Settings{
		ListenPort: 8083,
		ConnString: "server=mssql.au.ds.network; user id=service_user; password=Welcome100!; port=1433; database=arbitrag_articles", // "server=localhost; user id=sa; password=Welcome100; database=Articles",
		QryTimeout: 30,
		Workers:    10,
	}
}

// NewApp will return the app object for the article-api with cli parameters intergrated
func NewApp(settings *service.Settings) *cli.App {
	appSettingsDefault := DefaultSettings()
	serviceSettingsDefault := DefaultServiceSettings()
	app := cli.NewApp()

	app.Name = appSettingsDefault.Name
	app.Version = appSettingsDefault.Version
	app.Description = appSettingsDefault.Description
	app.Usage = appSettingsDefault.Usage

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "listen-port",
			Usage:       "HTTP port to listen on",
			EnvVar:      "APP_LISTEN_PORT",
			Destination: &settings.ListenPort,
			Value:       serviceSettingsDefault.ListenPort,
		},
		cli.StringFlag{
			Name:        "conn-string",
			Usage:       "The full db connection string for the mssql database",
			EnvVar:      "APP_CONN_STRING",
			Destination: &settings.ConnString,
			Value:       serviceSettingsDefault.ConnString,
		},
		cli.IntFlag{
			Name:        "qry-timeout",
			Usage:       "The query timeout in seconds",
			EnvVar:      "APP_QRY_TIMEOUT",
			Destination: &settings.QryTimeout,
			Value:       serviceSettingsDefault.QryTimeout,
		},
		cli.IntFlag{
			Name:        "workers",
			Usage:       "The number of database connection workers ",
			EnvVar:      "APP_WORKERS",
			Destination: &settings.Workers,
			Value:       serviceSettingsDefault.Workers,
		},
	}

	return app
}

// SetLogger sets the logger for this package
func SetLogger(logger *logrus.Logger) {
	log = logger
}
