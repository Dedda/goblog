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

func Index(articleProvider article.ArticleProvider, writer http.ResponseWriter, _ *http.Request) {
	categories, err := articleProvider.ListCategories()
	if err != nil {
		responseErr(err, writer)
		return
	}
	data, err := render(indexTemplate, &indexPage{
		Categories: categories,
	})
	if err != nil {
		responseErr(err, writer)
		return
	}
	responseOk(wrapContents(data), writer)
}

type indexPage struct {
	Categories []*article.ArticleCategory
}

func (*indexPage) TemplateText() string {
	return index_html
}

func ArticleList(articleProvider article.ArticleProvider, writer http.ResponseWriter, _ *http.Request, category string) {
	articles, err := articleProvider.ListArticles(category)
	if err != nil {
		responseErr(err, writer)
		return
	}
	data, err := render(articleListTemplate, &articleListPage{
		Category: category,
		Articles: articles,
	})
	if err != nil {
		responseErr(err, writer)
		return
	}
	responseOk(wrapContents(data), writer)
}

type articleListPage struct {
	Category string
	Articles []*article.ArticleMetaInfo
}

func (*articleListPage) TemplateText() string {
	return article_list_html
}

func Article(articleProvider article.ArticleProvider, writer http.ResponseWriter, category, id string) {
	a, err := articleProvider.RenderArticle(category, id)
	if err != nil {
		responseErr(err, writer)
		return
	}
	page := articlePage{
		Category: category,
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
	Category string
	Meta     *article.ArticleMetaInfo
	Rendered template.HTML
}

func (*articlePage) TemplateText() string {
	return article_html
}
