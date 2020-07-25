package main

import (
	"bytes"
	"github.com/fwchen/saury/repository"
	"github.com/labstack/echo/v4"
	"html/template"
	"log"
	"net/http"
)

func main() {
	tmp, err := template.ParseFiles("template/tmpl.html")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		appDb, err := repository.Connect("mongodb://localhost:27017", "saury")
		if err != nil {
			return err
		}
		galleryRepo := repository.NewGalleryRepository(appDb)
		galleries, err := galleryRepo.FindAll(20, 0)
		if err != nil {
			return err
		}

		var tpl bytes.Buffer
		if err := tmp.Execute(&tpl, galleries); err != nil {
			log.Fatal(err)
		}
		htmlResponse := tpl.String()

		return c.HTML(http.StatusOK, htmlResponse)
	})

	e.Logger.Fatal(e.Start(":1324"))
}
