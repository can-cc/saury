package main

import (
	"github.com/fwchen/saury/render"
	"github.com/fwchen/saury/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {

	e := echo.New()
	appDb, err := repository.Connect("mongodb://localhost:27017", "saury")
	if err != nil {
		log.Fatal(err)
	}
	galleryRepo := repository.NewGalleryRepository(appDb)

	e.GET("/", func(c echo.Context) error {

		galleries, err := galleryRepo.FindAll(20, 0)
		if err != nil {
			return err
		}

		htmlResponse, err := render.ParseIndex(galleries)
		if err != nil {
			return err
		}

		return c.HTML(http.StatusOK, htmlResponse)
	})

	e.GET("/album/:name/photo/:name", func(c echo.Context) error {

		values := c.ParamValues()
		albumName := values[0]
		photoUrl := values[1]

		galleries, err := galleryRepo.FindAll(20, 0)
		if err != nil {
			return err
		}

		htmlResponse, err := render.ParsePhoto(galleries, albumName, photoUrl)

		if err != nil {
			return err
		}

		return c.HTML(http.StatusOK, htmlResponse)
	})

	e.Static("/galleries", "galleries")

	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":1324"))
}
