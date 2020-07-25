package render

import (
	"bytes"
	"github.com/fwchen/saury/model"
	"html/template"
	"log"
)

func ParseFile(galleries []model.Album) string {
	tmp, err := template.ParseFiles("render/template/tmpl.html")
	if err != nil {
		log.Fatal(err)
	}
	var tpl bytes.Buffer
	if err := tmp.Execute(&tpl, galleries); err != nil {
		log.Fatal(err)
	}
	return tpl.String()
}
