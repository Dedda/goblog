package article

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"os"
	"sync"
	"time"
)

type articleRenderCache struct {
	articleFiles map[string]renderedArticleInfo
	mutex        sync.RWMutex
}

type renderedArticleInfo struct {
	lastModified time.Time
	data         []byte
}

func newArticleRenderCache() *articleRenderCache {
	return &articleRenderCache{
		articleFiles: make(map[string]renderedArticleInfo),
		mutex:        sync.RWMutex{},
	}
}

func (c *articleRenderCache) renderAndGet(filePath string) (*[]byte, error) {
	found, ok := c.getCached(filePath)
	if !ok {
		found, err := c.readFromDisk(filePath)
		if err != nil {
			return nil, err
		}
		return &found.data, err
	}
	fileStat, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	lastModified := fileStat.ModTime()
	if lastModified.After(found.lastModified) {
		found, err = c.readFromDisk(filePath)
		if err != nil {
			return nil, err
		}
	}
	return &found.data, err
}

func (c *articleRenderCache) getCached(filePath string) (renderedArticleInfo, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	article, ok := c.articleFiles[filePath]
	return article, ok
}

func (c *articleRenderCache) readFromDisk(filePath string) (renderedArticleInfo, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	fileStat, err := os.Stat(filePath)
	if err != nil {
		return renderedArticleInfo{}, err
	}
	lastModified := fileStat.ModTime()
	data, err := os.ReadFile(filePath)
	if err != nil {
		return renderedArticleInfo{}, err
	}
	rendered := renderedArticleInfo{
		lastModified: lastModified,
		data:         c.render(&data),
	}
	c.articleFiles[filePath] = rendered
	return rendered, nil
}

func (c *articleRenderCache) render(data *[]byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(*data)
	htmlFlags := html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	return markdown.Render(doc, renderer)
}
