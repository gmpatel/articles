package controller

import (
	"net/http"

	"github.com/gmpatel/articles"
	"github.com/gmpatel/articles/model"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//GetArticles returns
func GetArticles(c *gin.Context, repository articles.Repository) {
	// swagger:operation GET /healthz system health
	//
	// Health Check
	// ---
	// produces:
	// - application/json
	// responses:
	//  '200':
	//    description: "status ok"
	//    type: string
	//  default:
	//    description: unexpected error
	//    schema:
	//      "$ref": "#/responses/DefaultError"

	id := c.Param("id")

	c.String(http.StatusOK, id)
}

//PostArticles returns
func PostArticles(c *gin.Context, repository articles.Repository) {
	// swagger:operation POST /articles article postArticle
	//
	// Saves the articles
	// ---
	// Consumes:
	//  - application/json
	// produces:
	// - application/json
	// responses:
	//  '200':
	//    description: "ok"
	//    type: string
	//  default:
	//    description: unexpected error
	//    schema:
	//     "$ref": "#/responses/DefaultError"

	var articleRequest model.ArticleRequest
	if err := c.ShouldBindWith(&articleRequest, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

//Health returns the health of the service
func Health(c *gin.Context, repository articles.Repository) {
	// swagger:operation GET /healthz system health
	//
	// Health Check
	// ---
	// produces:
	// - application/json
	// responses:
	//  '200':
	//    description: "status ok"
	//    type: string
	//  default:
	//    description: unexpected error
	//    schema:
	//      "$ref": "#/responses/DefaultError"

	c.String(http.StatusOK, "ok")
}

// Ready returns the liveness of service
func Ready(c *gin.Context, repository articles.Repository) {
	// swagger:operation GET /readiness system ready
	//
	// Ready Check
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//  '200':
	//    description: "status ok"
	//    type: string
	//  default:
	//    description: unexpected error
	//    schema:
	//       "$ref": "#/responses/DefaultError"
	c.String(http.StatusOK, "ok")
}

// SetLogger sets the logger for this package
func SetLogger(logger *logrus.Logger) {
	log = logger
}
