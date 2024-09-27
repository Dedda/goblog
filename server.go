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
	serveMux.HandleFunc("GET /assets/style.css", assets.StyleCSS)
	server := &http.Server{
		Handler: serveMux,
		Addr:    address,
	}
	return server.ListenAndServe()
}

func articleList(writer http.ResponseWriter, request *http.Request) {
	pages.ArticleList(articleProvider, writer, request)
}
