package model

import (
	"time"
)

// ErrorResponse ...
type ErrorResponse struct {
	// the Code of the error
	// required: false
	ErrorCode string `json:"errorCode"`

	// the Message of the error
	// required: true
	ErrorMessage string `json:"errorMessage"`
}

// A DefaultError is an error that is used when the required input fails validation or an unexpected internal error occurs.
// swagger:response DefaultError
type DefaultError struct {
	// The error message
	// in: body
	Body ErrorResponse
}

// ArticleRequest defines the structure of the article API request body
// swagger:model
type ArticleRequest struct {
	// the ID of the article
	// required: true
	ID int64 `json:"id" binding:"required"`

	// the Title of the article
	// required: true
	Title string `json:"title" binding:"required"`

	// the Date of the article
	// required: true
	Date time.Time `json:"date" binding:"required"`

	// the Body of the article
	// required: true
	Body string `json:"body" binding:"required"`

	// the Tags of the article
	// required: true
	Tags []string `json:"tags" binding:"required"`
}
