package endpoint

import (
	"fmt"
	"net/http"
	"time"

	"github.com/articles"
	"github.com/sirupsen/logrus"

	"github.com/articles/controller"
	"github.com/gin-gonic/gin"
)

var (
	log *logrus.Logger
)

func init() {
	SetLogger(logrus.StandardLogger())
}

func setupRouter(repository articles.Repository) *gin.Engine {
	router := getRouter()

	router.POST("/articles",
		func(context *gin.Context) {
			controller.PostArticles(context, repository)
		})

	router.GET("/articles",
		func(context *gin.Context) {
			controller.GetArticles(context, repository)
		})

	router.GET("/articles/:id",
		func(context *gin.Context) {
			controller.GetArticles(context, repository)
		})

	router.GET("/healthz",
		func(context *gin.Context) {
			controller.Health(context, repository)
		})

	router.GET("/readiness",
		func(context *gin.Context) {
			controller.Ready(context, repository)
		})

	router.GET("/servertime",
		func(context *gin.Context) {
			context.String(http.StatusOK, time.Now().UTC().String())
		})

	return router
}

func getRouter() *gin.Engine {
	//Set release mode of libraries
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.RedirectFixedPath = true

	//Add Middlewares
	router.Use(gin.Recovery())
	router.Use(ginLogMiddleware())
	router.Use(ginResponseHeadersMiddleware())

	//Return
	return router
}

func ginResponseHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Served-By", "GP's Articles API Service")
	}
}

func ginLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.RequestURI()

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		comment := c.Errors.String()
		userAgent := c.Request.UserAgent()

		logs := fmt.Sprintln("method :", method, "|",
			"path :", path, "|",
			"latency :", latency, "|",
			"ip :", clientIP, "|",
			"comment :", comment, "|",
			"statusCode :", statusCode, "|",
			"user-agent :", userAgent,
		)

		if statusCode > 399 {
			log.Error(fmt.Errorf("Error returned from the GIN/HTTP context with the status code %d", statusCode), logs, "")
		} else {
			log.Info(string(statusCode), logs, "")
		}
	}
}
