package article

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileSystemArticleProvider struct {
	directory string
	articles  []ArticleMetaInfo
	cache     *articleRenderCache
}

func NewFileSystemArticleProvider(directory string) (*FileSystemArticleProvider, error) {
	var meta []ArticleMetaInfo
	raw, err := os.ReadFile(directory + "/articles.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(raw, &meta)
	for i, _ := range meta {
		m := &meta[i]
		m.mdFilename = fmt.Sprintf("%s/%s.md", directory, m.Id)
	}
	fmt.Printf("Found %d article entries.", len(meta))
	provider := FileSystemArticleProvider{
		directory: directory,
		articles:  meta,
		cache:     newArticleRenderCache(),
	}
	return &provider, err
}

func (f FileSystemArticleProvider) ListArticles() ([]*ArticleMetaInfo, error) {
	articles := make([]*ArticleMetaInfo, len(f.articles))
	for i, meta := range f.articles {
		articles[i] = &meta
	}
	return articles, nil
}

func (f FileSystemArticleProvider) GetArticle(id string) (*ArticleMetaInfo, error) {
	articles, err := f.ListArticles()
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

func (f FileSystemArticleProvider) RenderArticle(id string) (RenderedArticle, error) {
	meta, err := f.GetArticle(id)
	if err != nil {
		return RenderedArticle{}, err
	}
	if meta == nil {
		return RenderedArticle{}, errors.New("Cannot find article " + id)
	}
	rendered, err := f.cache.renderAndGet(meta.mdFilename)
	if err != nil {
		return RenderedArticle{}, err
	}
	return RenderedArticle{
		Meta:     *meta,
		Rendered: rendered,
	}, err
}
