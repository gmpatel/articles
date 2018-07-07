package controller

import (
	"net/http"
	"strconv"

	"github.com/gmpatel/articles"
	"github.com/gmpatel/articles/model"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//GetTags returns
func GetTags(c *gin.Context, repository articles.Repository) {
	param1 := c.Param("tagName")
	param2 := c.Param("date")

	data, err := repository.GetTag(param1, param2)
	if err != nil {
		c.JSON(
			http.StatusBadRequest, model.ErrorResponse{ErrorMessage: err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK, data,
	)
	return
}

//GetArticles returns
func GetArticles(c *gin.Context, repository articles.Repository) {
	var id int64

	param1 := c.Param("id")
	if len(param1) > 0 {
		val, err := strconv.ParseInt(param1, 10, 64)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError, model.ErrorResponse{ErrorMessage: err.Error()},
			)
			return
		}
		id = val
	}

	data, err := repository.GetArticles(id)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, model.ErrorResponse{ErrorMessage: err.Error()},
		)
		return
	}

	if id > 0 {
		if len(data) > 0 {
			c.JSON(
				http.StatusOK, data[0],
			)
		} else {
			c.JSON(
				http.StatusNotFound, nil,
			)
		}

	} else {
		c.JSON(
			http.StatusOK, data,
		)
	}

	return
}

//PostArticle returns
func PostArticle(c *gin.Context, repository articles.Repository) {
	var ArticleModel model.ArticleModel
	if err := c.ShouldBindWith(&ArticleModel, binding.JSON); err != nil {
		c.JSON(
			http.StatusBadRequest, model.ErrorResponse{ErrorMessage: err.Error()},
		)
		return
	}

	id, err := repository.StoreArticle(&ArticleModel)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, model.ErrorResponse{ErrorMessage: err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		id,
	)
}

//Health returns the health of the service
func Health(c *gin.Context, repository articles.Repository) {

	// Can do more here depends on the service dependencies

	c.String(
		http.StatusOK,
		"OK",
	)
}

// Ready returns the liveness of service
func Ready(c *gin.Context, repository articles.Repository) {

	// Can do more here depends on the service dependencies

	c.String(
		http.StatusOK,
		"OK",
	)
}

// SetLogger sets the logger for this package
func SetLogger(logger *logrus.Logger) {
	log = logger
}
