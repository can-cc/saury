package main

import (
	"github.com/fwchen/saury/render"
	"github.com/fwchen/saury/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {

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

		htmlResponse, err := render.ParseFile(galleries)
		if err != nil {
			return err
		}

		return c.HTML(http.StatusOK, htmlResponse)
	})

	e.Static("/*", "")

	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":1324"))
}
