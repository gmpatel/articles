package articles

import "github.com/gmpatel/articles/model"

// Repository inteface
type Repository interface {
	StoreArticle(article *model.ArticleModel) (int64, *string, error)
	GetArticles(id int64) ([]model.ArticleModel, error)
	GetTag(name string, date string) (*model.TagModel, error)
}

// Service inteface
type Service interface {
	Start()
	Stop()
}
