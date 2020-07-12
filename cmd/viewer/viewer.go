package main

import (
	"github.com/fwchen/saury/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)



func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		queryPage := c.QueryParam("page")
		page, err := strconv.Atoi(queryPage)
		if err != nil {
			page = 0
		}
		appDb, err := repository.Connect("mongodb://localhost:27017", "saury")
		if err != nil {
			return err
		}
		galleryRepo := repository.NewGalleryRepository(appDb)
		galleries, err := galleryRepo.FindAll(20, int64(page))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, galleries)
	})


	e.Logger.Fatal(e.Start(":1323"))
}