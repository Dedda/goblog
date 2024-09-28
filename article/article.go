package article

import (
	"time"
)

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
	ListArticles() ([]*ArticleMetaInfo, error)
	GetArticle(id string) (*ArticleMetaInfo, error)
	RenderArticle(id string) (RenderedArticle, error)
}
