package render

import (
	"bytes"
	"github.com/fwchen/saury/model"
	"html/template"
)

func ParseIndex(galleries []model.Album) (string, error) {
	tmp, err := template.ParseFiles("template/index.html")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := tmp.Execute(&tpl, galleries); err != nil {
		return "", err
	}
	return tpl.String(), nil
}

func ParsePhoto(galleries []model.Album, alumName string, photoName string, prevPhoto string, nextPhoto string) (string, error) {
	tmp, err := template.ParseFiles("template/photo.html")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := tmp.Execute(&tpl, map[string]interface{}{
		"galleries": galleries,
		"alumName":  alumName,
		"photoName": photoName,
		"prevPhoto": prevPhoto,
		"nextPhoto": nextPhoto,
	}); err != nil {
		return "", err
	}
	return tpl.String(), nil
}
