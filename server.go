package main

import (
	"bytes"
	"github.com/Dedda/goblog/pages"
	"log"
	"net/http"
)

func startServer(address string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/articles/{$}", reqArticleList)
	server := &http.Server{
		Handler: serveMux,
		Addr:    address,
	}
	return server.ListenAndServe()
}

func reqArticleList(writer http.ResponseWriter, request *http.Request) {
	buf := bytes.Buffer{}
	articles, err := articleProvider.ListArticles()
	if err != nil {
		log.Fatal(err)
	}
	err = pages.ArticleListTemplate.Render(&buf, &pages.ArticleListPage{
		Articles: articles,
	})
	if err != nil {
		log.Fatal(err)
	}
	_, err = writer.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
}
