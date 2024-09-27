package pages

import (
	"bytes"
	"github.com/tylermmorton/tmpl"
)

func render[T tmpl.TemplateProvider](template tmpl.Template[T], data T) ([]byte, error) {
	buf := bytes.Buffer{}
	err := template.Render(&buf, data)
	return buf.Bytes(), err
}
