package render

import (
	"bytes"
	"github.com/fwchen/saury/model"
	"html/template"
)

func ParseFile(galleries []model.Album) (string, error) {
	tmp, err := template.ParseFiles("template/tmpl.html")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := tmp.Execute(&tpl, galleries); err != nil {
		return "", err
	}
	return tpl.String(), nil
}
