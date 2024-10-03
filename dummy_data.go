package main

import (
	"github.com/Dedda/goblog/article"
	"time"
)

type dummyArticleProvider struct{}

func (d dummyArticleProvider) ListArticles() ([]*article.ArticleMetaInfo, error) {
	articles := []*article.ArticleMetaInfo{
		{
			Id:    "1",
			Title: "Article 1",
			Created: article.Date{
				Time: time.Now(),
			},
		},
		{
			Id:    "2",
			Title: "Article 2",
			Created: article.Date{
				Time: time.Now(),
			},
		},
	}
	return articles, nil
}

func (d dummyArticleProvider) GetArticle(id string) (*article.ArticleMetaInfo, error) {
	articles, err := d.ListArticles()
	if err != nil {
		return nil, err
	}
	for _, a := range articles {
		if a.Id == id {
			return a, nil
		}
	}
	return nil, nil
}

func (d dummyArticleProvider) RenderArticle(id string) (article.RenderedArticle, error) {
	meta, err := d.GetArticle(id)
	if err != nil {
		return article.RenderedArticle{}, err
	}
	switch id {
	case "1":
		data := []byte("<h1>Article 1</h1>")
		return article.RenderedArticle{
			Meta:     *meta,
			Rendered: &data,
		}, nil
	case "2":
		data := []byte("<h1>Article 2</h1>")
		return article.RenderedArticle{
			Meta:     *meta,
			Rendered: &data,
		}, nil
	default:
		return article.RenderedArticle{}, nil
	}
}
