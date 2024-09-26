package main

import (
	"github.com/Dedda/goblog/article"
	"log"
	"os"
)

var (
	articleProvider article.ArticleProvider
)

func main() {
	articleProvider = &dummyArticleProvider{}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := startServer(":" + port); err != nil {
		log.Fatal(err)
	}
}
