package article

import (
	"time"
)

type ArticleMetaInfo struct {
	Id         string
	Title      string
	Created    time.Time
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

type CachedArticleProvider struct {
	cache *articleRenderCache
}

func NewCachedArticleProvider() *CachedArticleProvider {
	return &CachedArticleProvider{
		cache: newArticleRenderCache(),
	}
}

func (p *CachedArticleProvider) ListArticles() ([]*ArticleMetaInfo, error) {
	return []*ArticleMetaInfo{}, nil
}

func (p *CachedArticleProvider) GetArticle(id string) (*ArticleMetaInfo, error) {
	return nil, nil
}

func (p *CachedArticleProvider) RenderArticle(id string) (RenderedArticle, error) {
	meta, err := p.GetArticle(id)
	if err != nil {
		return RenderedArticle{}, err
	}
	rendered, err := p.cache.renderAndGet(meta.mdFilename)
	if err != nil {
		return RenderedArticle{}, err
	}
	return RenderedArticle{
		Meta:     *meta,
		Rendered: rendered,
	}, err
}
