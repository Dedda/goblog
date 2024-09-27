package pages

import (
	_ "embed"
	"github.com/tylermmorton/tmpl"
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

	//go:embed templates/article.tmpl.html
	article_html    string
	articleTemplate = tmpl.MustCompile(&articlePage{})
)
