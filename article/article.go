package article

import (
	"time"
)

type ArticleCategory struct {
	Id       string            `json:"id"`
	Title    string            `json:"title"`
	Articles []ArticleMetaInfo `json:"articles"`
}

type ArticleMetaInfo struct {
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	Created    time.Time `json:"created,omitempty"`
	mdFilename string
}

type RenderedArticle struct {
	Meta     ArticleMetaInfo
	Rendered *[]byte
}

type ArticleProvider interface {
	ListCategories() ([]*ArticleCategory, error)
	ListArticles(category string) ([]*ArticleMetaInfo, error)
	GetArticle(category, id string) (*ArticleMetaInfo, error)
	RenderArticle(category, id string) (RenderedArticle, error)
}
