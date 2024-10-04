package assets

import (
	_ "embed"
	"log"
	"net/http"
)

var (
	//go:embed w3-blog.css
	w3BlogCss []byte

	//go:embed extras.css
	extrasCss []byte

	//go:embed chroma-github.css
	syntaxGithubCss []byte
)

func StyleCSS(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/css; charset=utf-8")
	_, err := writer.Write(w3BlogCss)
	if err != nil {
		log.Fatal(err)
	}
}

func ExtrasCSS(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/css; charset=utf-8")
	_, err := writer.Write(extrasCss)
	if err != nil {
		log.Fatal(err)
	}
}

func SyntaxGithubCSS(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/css; charset=utf-8")
	_, err := writer.Write(syntaxGithubCss)
	if err != nil {
		log.Fatal(err)
	}
}
