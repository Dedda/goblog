package article

import (
	"time"
)

type Date struct {
	Time time.Time
}

func (d *Date) UnmarshalJSON(data []byte) error {
	date, err := time.Parse(`"02.01.2006"`, string(data))
	if err != nil {
		return err
	}
	d.Time = date
	return nil
}

type ArticleCategory struct {
	Id       string            `json:"id"`
	Title    string            `json:"title"`
	Articles []ArticleMetaInfo `json:"articles"`
}

type ArticleMetaInfo struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Created    Date   `json:"created,omitempty"`
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
