package model

import (
	"time"
)

// ErrorResponse ...
type ErrorResponse struct {
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// A DefaultError is an error that is used when the required input fails validation or an unexpected internal error occurs.
type DefaultError struct {
	Body ErrorResponse
}

// ArticleModel defines the structure of the article API request body
type ArticleModel struct {
	ID         int64     `json:"id,omitempty"`
	Title      string    `json:"title" binding:"required"`
	DateTime   time.Time `json:"-"`
	Date       string    `json:"date,omitempty"`
	Body       string    `json:"body" binding:"required"`
	Tags       []string  `json:"tags" binding:"required"`
	TagsString string    `json:"-"`
}

// TagModel defines the structure of the tag API response object
type TagModel struct {
	ID                int64    `json:"-"`
	Tag               string   `json:"tag,omitempty"`
	ArticlesString    string   `json:"-"`
	RelatedTagsString string   `json:"-"`
	Count             int      `json:"count,omitempty"`
	Articles          []string `json:"articles,omitempty"`
	RelatedTags       []string `json:"related_tags,omitempty"`
}
