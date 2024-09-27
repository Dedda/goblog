package pages

import (
	"bytes"
	_ "embed"
	"github.com/Dedda/goblog/article"
	"github.com/tylermmorton/tmpl"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed templates/wrapper.tmpl.html
	wrapper_html    string
	wrapperTemplate = tmpl.MustCompile(&wrapperPage{})

	//go:embed templates/index.tmpl.html
	index_html    string
	indexTemplate = tmpl.MustCompile(&indexPage{})

	//go:embed templates/article_list.tmpl.html
	article_list_html   string
	articleListTemplate = tmpl.MustCompile(&articleListPage{})
)

func wrapContents(contents []byte) []byte {
	buf := bytes.Buffer{}
	wrapper := wrapperPage{
		Contents: template.HTML(contents),
	}
	err := wrapperTemplate.Render(&buf, &wrapper)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

type wrapperPage struct {
	Contents template.HTML
}

func (*wrapperPage) TemplateText() string {
	return wrapper_html
}

func Index(writer http.ResponseWriter, _ *http.Request) {
	buf := bytes.Buffer{}
	err := indexTemplate.Render(&buf, &indexPage{})
	if err != nil {
		log.Fatal(err)
	}
	_, err = writer.Write(wrapContents(buf.Bytes()))
	if err != nil {
		log.Fatal(err)
	}
	writer.Header().Add("Content-Type", "text/html; charset=utf-8")
}

type indexPage struct{}

func (*indexPage) TemplateText() string {
	return index_html
}

func ArticleList(articleProvider article.ArticleProvider, writer http.ResponseWriter, _ *http.Request) {
	articles, err := articleProvider.ListArticles()
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.Buffer{}
	err = articleListTemplate.Render(&buf, &articleListPage{
		Articles: articles,
	})
	if err != nil {
		log.Fatal(err)
	}
	_, err = writer.Write(wrapContents(buf.Bytes()))
	if err != nil {
		log.Fatal(err)
	}
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
}

type articleListPage struct {
	Articles []*article.ArticleMetaInfo
}

func (*articleListPage) TemplateText() string {
	return article_list_html
}
