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
	"time"
)

var (
	htmlFormatter  *html.Formatter
	highlightStyle *chroma.Style
)

func init() {
	htmlFormatter = html.New(html.WithClasses(true), html.TabWidth(2))
	if htmlFormatter == nil {
		panic("couldn't create html formatter")
	}
	highlightStyle = styles.GitHub
}

func renderMD(data *[]byte) []byte {
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
