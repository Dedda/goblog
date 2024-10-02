package main

import (
	"github.com/Dedda/goblog/assets"
	"github.com/Dedda/goblog/pages"
	"net/http"
)

func startServer(address string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("GET /{$}", pages.Index)
	serveMux.HandleFunc("GET /articles/{$}", articleList)
	serveMux.HandleFunc("GET /article/{id}", articlePage)
	serveMux.HandleFunc("GET /assets/style.css", assets.StyleCSS)
	serveMux.HandleFunc("GET /assets/extras.css", assets.ExtrasCSS)
	server := &http.Server{
		Handler: serveMux,
		Addr:    address,
	}
	return server.ListenAndServe()
}

func articleList(writer http.ResponseWriter, request *http.Request) {
	pages.ArticleList(articleProvider, writer, request)
}

func articlePage(writer http.ResponseWriter, request *http.Request) {
	pages.Article(articleProvider, writer, request.PathValue("id"))
}
