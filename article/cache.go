package article

import (
	"fmt"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	mdhtml "github.com/gomarkdown/markdown/html"
	"io"
	"os"
	"sync"
	"time"
)

var (
	htmlFormatter  *html.Formatter
	highlightStyle *chroma.Style
)

type articleRenderCache struct {
	articleFiles map[string]renderedArticleInfo
	mutex        sync.RWMutex
}

type renderedArticleInfo struct {
	lastModified time.Time
	data         []byte
}

func init() {
	htmlFormatter = html.New(html.WithClasses(true), html.TabWidth(2))
	if htmlFormatter == nil {
		panic("couldn't create html formatter")
	}
	highlightStyle = styles.GitHub
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
	//extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	//p := parser.NewWithExtensions(extensions)
	renderer := newCustomizedRenderer()
	return markdown.ToHTML(*data, nil, renderer)
}

func newCustomizedRenderer() *mdhtml.Renderer {
	opts := mdhtml.RendererOptions{
		Flags:          mdhtml.CommonFlags,
		RenderNodeHook: myRenderHook,
	}
	return mdhtml.NewRenderer(opts)
}

func myRenderHook(writer io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if code, ok := node.(*ast.CodeBlock); ok {
		if err := renderCode(writer, code, entering); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "[ERROR] %s: %s\n", time.Now().Format(time.RFC3339), err)
		}
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}

func renderCode(writer io.Writer, codeBlock *ast.CodeBlock, entering bool) error {
	lang := string(codeBlock.Info)
	return htmlHighlight(writer, string(codeBlock.Literal), lang, "go")
}

func htmlHighlight(writer io.Writer, codeBlock, lang, defaultLang string) error {
	if lang == "" {
		lang = defaultLang
	}
	l := lexers.Get(lang)
	if l == nil {
		l = lexers.Analyse(codeBlock)
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	it, err := l.Tokenise(nil, codeBlock)
	if err != nil {
		return err
	}
	return htmlFormatter.Format(writer, highlightStyle, it)
}
