package article

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileSystemArticleProvider struct {
	directory  string
	categories []ArticleCategory
	cache      *articleRenderCache
}

func NewFileSystemArticleProvider(directory string) (*FileSystemArticleProvider, error) {
	var categories []ArticleCategory
	raw, err := os.ReadFile(directory + "/articles.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(raw, &categories)
	articles := 0
	for _, category := range categories {
		for i, _ := range category.Articles {
			m := &category.Articles[i]
			m.mdFilename = fmt.Sprintf("%s/%s/%s.md", directory, category.Id, m.Id)
			articles += 1
		}
	}
	fmt.Printf("Found %d article entries in %d categories.", articles, len(categories))
	provider := FileSystemArticleProvider{
		directory:  directory,
		categories: categories,
		cache:      newArticleRenderCache(),
	}
	return &provider, err
}

func (f *FileSystemArticleProvider) getCategory(id string) (*ArticleCategory, error) {
	for _, category := range f.categories {
		if category.Id == id {
			return &category, nil
		}
	}
	return nil, errors.New("Category not found")
}

func (f *FileSystemArticleProvider) ListCategories() ([]*ArticleCategory, error) {
	categories := make([]*ArticleCategory, len(f.categories))
	for i, c := range f.categories {
		categories[i] = &c
	}
	return categories, nil
}

func (f *FileSystemArticleProvider) ListArticles(category string) ([]*ArticleMetaInfo, error) {
	c, err := f.getCategory(category)
	if err != nil {
		return nil, err
	}
	articles := make([]*ArticleMetaInfo, len(c.Articles))
	for i, meta := range c.Articles {
		articles[i] = &meta
	}
	return articles, nil
}

func (f *FileSystemArticleProvider) GetArticle(category, id string) (*ArticleMetaInfo, error) {
	articles, err := f.ListArticles(category)
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

func (f *FileSystemArticleProvider) RenderArticle(category, id string) (RenderedArticle, error) {
	meta, err := f.GetArticle(category, id)
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
