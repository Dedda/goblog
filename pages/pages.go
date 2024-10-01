package pages

import (
	_ "embed"
	"fmt"
	"github.com/Dedda/goblog/article"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func responseErr(err error, writer http.ResponseWriter) {
	_, e := fmt.Fprintf(os.Stderr, "[ERROR] %s: %s", time.Now().Format(time.RFC3339), err)
	if e != nil {
		log.Fatal(e)
	}
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.WriteHeader(http.StatusInternalServerError)
}

func responseOk(data []byte, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := writer.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func wrapContents(contents []byte) []byte {
	wrapper := wrapperPage{
		Contents: template.HTML(contents),
	}
	data, err := render(wrapperTemplate, &wrapper)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

type wrapperPage struct {
	Contents template.HTML
}

func (*wrapperPage) TemplateText() string {
	return wrapper_html
}

func Index(writer http.ResponseWriter, _ *http.Request) {
	data, err := render(indexTemplate, &indexPage{})
	if err != nil {
		responseErr(err, writer)
		return
	}
	responseOk(wrapContents(data), writer)
}

type indexPage struct{}

func (*indexPage) TemplateText() string {
	return index_html
}

func ArticleList(articleProvider article.ArticleProvider, writer http.ResponseWriter, _ *http.Request) {
	articles, err := articleProvider.ListArticles()
	if err != nil {
		responseErr(err, writer)
		return
	}
	data, err := render(articleListTemplate, &articleListPage{
		Articles: articles,
	})
	if err != nil {
		responseErr(err, writer)
		return
	}
	responseOk(wrapContents(data), writer)
}

type articleListPage struct {
	Articles []*article.ArticleMetaInfo
}

func (*articleListPage) TemplateText() string {
	return article_list_html
}

func Article(articleProvider article.ArticleProvider, writer http.ResponseWriter, id string) {
	a, err := articleProvider.RenderArticle(id)
	if err != nil {
		responseErr(err, writer)
		return
	}
	page := articlePage{
		Meta:     &a.Meta,
		Rendered: template.HTML(*a.Rendered),
	}
	data, err := render(articleTemplate, &page)
	if err != nil {
		responseErr(err, writer)
		return
	}
	responseOk(wrapContents(data), writer)
}

type articlePage struct {
	Meta     *article.ArticleMetaInfo
	Rendered template.HTML
}

func (*articlePage) TemplateText() string {
	return article_html
}
