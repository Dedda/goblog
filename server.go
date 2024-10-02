package main

import (
	"github.com/Dedda/goblog/assets"
	"github.com/Dedda/goblog/pages"
	"net/http"
)

func startServer(address string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("GET /{$}", indexPage)
	serveMux.HandleFunc("GET /articles/{category}/{$}", articleList)
	serveMux.HandleFunc("GET /article/{category}/{id}", articlePage)
	serveMux.HandleFunc("GET /assets/style.css", assets.StyleCSS)
	serveMux.HandleFunc("GET /assets/extras.css", assets.ExtrasCSS)
	server := &http.Server{
		Handler: serveMux,
		Addr:    address,
	}
	return server.ListenAndServe()
}

func indexPage(writer http.ResponseWriter, request *http.Request) {
	pages.Index(articleProvider, writer, request)
}

func articleList(writer http.ResponseWriter, request *http.Request) {
	category := request.PathValue("category")
	pages.ArticleList(articleProvider, writer, request, category)
}

func articlePage(writer http.ResponseWriter, request *http.Request) {
	category := request.PathValue("category")
	id := request.PathValue("id")
	pages.Article(articleProvider, writer, category, id)
}
