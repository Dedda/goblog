package pages

import (
	_ "embed"
	"github.com/Dedda/goblog/article"
	"github.com/tylermmorton/tmpl"
)

var (
	//go:embed templates/article_list.tmpl.html
	article_list_html   string
	ArticleListTemplate = tmpl.MustCompile(&ArticleListPage{})
)

type ArticleListPage struct {
	Articles []*article.ArticleMetaInfo
}

func (*ArticleListPage) TemplateText() string {
	return article_list_html
}
