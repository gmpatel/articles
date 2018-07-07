package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gmpatel/articles/model"

	_ "github.com/denisenkom/go-mssqldb" // Blank import is important to register the sql drivers "mssql" / "sqlserver"
	"github.com/sirupsen/logrus"
)

// SetLogger sets the logger for this package
func SetLogger(logger *logrus.Logger) {
	log = logger
}

// SQLRepository object to provide repository methods talking to the db
type SQLRepository struct {
	db      *sql.DB
	mutex   *sync.Mutex
	timeout time.Duration
}

// NewRepository returns the repository object constructed with the parameter given
func NewRepository(connString string, conns int, timeoutMillis int) (*SQLRepository, error) {
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(conns)
	db.SetMaxIdleConns(conns)

	return &SQLRepository{
		db:      db,
		mutex:   new(sync.Mutex),
		timeout: time.Millisecond * time.Duration(timeoutMillis),
	}, nil
}

// Start will start the repository
func (repo *SQLRepository) start() error {
	return repo.ensureConnected()
}

// Stop closes any resources used by the repository
func (repo *SQLRepository) stop() {
	err := repo.db.Close()
	if err != nil {
		log.Debugf("Failed to close connection: %v", err)
	}
}

// StoreArticle stores the given article into the database
func (repo *SQLRepository) StoreArticle(article *model.ArticleModel) (int64, error) {
	storedProc := "[dbo].[spPostArticle]"
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	tags := strings.Join(article.Tags, ",")

	cmd := fmt.Sprintf("EXEC %s @title, @body, @tags", storedProc)
	rows, err := repo.db.QueryContext(ctx, cmd, sql.Named("title", article.Title), sql.Named("body", article.Body), sql.Named("tags", tags))
	if err != nil {
		log.Errorf("failed to execute stored procedure '%s': %v", storedProc, err)
		return 0, err
	}

	ret := getScallerValue(rows)
	return ret.(int64), err
}

// GetArticles stores the given article into the database
func (repo *SQLRepository) GetArticles(id int64) ([]model.ArticleModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	proc := "[dbo].[spGetArticles]"
	cmd := fmt.Sprintf("EXEC %s %d", proc, id)
	articles := make([]model.ArticleModel, 0)
	rows, _, err := repo.queryContext(ctx, cmd)

	if err != nil {
		log.Errorf("Failed to query for GetArticles(%d) method: %v", id, err)
		return nil, err
	}

	for rows.Next() {
		article := model.ArticleModel{}

		scanErr := rows.Scan(
			&article.ID,
			&article.Title,
			&article.DateTime,
			&article.Body,
			&article.TagsString,
		)
		article.Date = article.DateTime.Format("2006-01-02")
		article.Tags = strings.Split(article.TagsString, ",")

		if scanErr != nil {
			log.Errorf("Failed to scan row cells: %v", scanErr)
			return nil, scanErr
		}

		articles = append(articles, article)
	}

	return articles, nil
}

// GetTag returns the tag stats for the given date
func (repo *SQLRepository) GetTag(name string, date string) (*model.TagModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	proc := "[dbo].[spGetTags]"
	cmd := fmt.Sprintf("EXEC %s '%s', '%s'", proc, name, date)
	rows, _, err := repo.queryContext(ctx, cmd)

	if err != nil {
		log.Errorf("Failed to query for GetTag(%s, %s) method: %v", name, date)
		return nil, err
	}

	for rows.Next() {
		tag := model.TagModel{}

		scanErr := rows.Scan(
			&tag.ID,
			&tag.Tag,
			&tag.ArticlesString,
			&tag.RelatedTagsString,
		)
		tag.Articles = strings.Split(tag.ArticlesString, ",")
		tag.RelatedTags = strings.Split(tag.RelatedTagsString, ",")
		tag.Count = len(tag.Articles)

		if scanErr != nil {
			log.Errorf("Failed to scan row cells: %v", scanErr)
			return nil, scanErr
		}

		return &tag, nil
	}

	return nil, fmt.Errorf("No records found for tag '%s' for the date '%s'", name, date)
}
