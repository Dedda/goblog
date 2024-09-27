package assets

import (
	_ "embed"
	"log"
	"net/http"
)

var (
	//go:embed w3-blog.css
	W3BlogCss []byte
)

func StyleCSS(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/css; charset=utf-8")
	_, err := writer.Write(W3BlogCss)
	if err != nil {
		log.Fatal(err)
	}
}
