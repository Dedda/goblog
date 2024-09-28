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
	provider, err := article.NewFileSystemArticleProvider("./articles") //&dummyArticleProvider{}
	if err != nil {
		log.Fatal(err)
	}
	articleProvider = provider
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := startServer(":" + port); err != nil {
		log.Fatal(err)
	}
}
