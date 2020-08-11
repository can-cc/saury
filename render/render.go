package render

import (
	"bytes"
	"github.com/fwchen/saury/model"
	"html/template"
)

func ParseIndex(galleries []model.Album, albumName string, photos []model.Photo, currentPage int, pageCount int) (string, error) {

	tmp, err := template.New("Index").Funcs(template.FuncMap{
		"Increase": func(num int) int {
			return num + 1
		},
		"Reduce": func(num int) int {
			return num - 1
		},
		"MakeRange": func(min int, max int) []int {
			a := make([]int, max-min+1)
			for i := range a {
				a[i] = min + i
			}
			return a
		}}).ParseFiles("template/index.html")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := tmp.Execute(&tpl, map[string]interface{}{
		"galleries":   galleries,
		"albumName":   albumName,
		"photos":      photos,
		"currentPage": currentPage,
		"pageCount":   pageCount,
	}); err != nil {
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
