package main

import (
	"errors"
	"github.com/fwchen/saury/render"
	"github.com/fwchen/saury/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const (
	DEFAULT_PHOTO_LIMIT = 21
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

	e.GET("/album/:albumName", func(c echo.Context) error {

		galleries, err := galleryRepo.FindAll(20, 0)
		if err != nil {
			return err
		}

		albumName := c.Param("albumName")
		unescapeAlbumName, err := url.PathUnescape(albumName)
		if err != nil {
			return err
		}

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			page = 1
		}

		photos, err := galleryRepo.FindPhotos(unescapeAlbumName, DEFAULT_PHOTO_LIMIT, (page-1)*DEFAULT_PHOTO_LIMIT)
		if err != nil {
			return err
		}

		photosCount, err := galleryRepo.FindPhotosCount(unescapeAlbumName)
		if err != nil {
			return err
		}

		pageCount := int(math.Ceil(float64(photosCount) / DEFAULT_PHOTO_LIMIT))
		htmlResponse, err := render.ParseAlbum(galleries, unescapeAlbumName, photos, page, pageCount)
		if err != nil {
			return err
		}

		return c.HTML(http.StatusOK, htmlResponse)
	})

	e.GET("/album/:albumName/photo/:photoName", func(c echo.Context) error {

		albumName := c.Param("albumName")
		photoName := c.Param("photoName")

		galleries, err := galleryRepo.FindAll(20, 0)
		if err != nil {
			return err
		}

		unescapeAlbumName, err := url.PathUnescape(albumName)
		if err != nil {
			return err
		}

		album, err := galleryRepo.FindByName(unescapeAlbumName)
		if err != nil {
			return err
		}

		var targetPhoto string
		var prevPhoto string
		var nextPhoto string
		unescapePhotoName, err := url.PathUnescape(photoName)
		if err != nil {
			return err
		}

		photosLen := len(album.Photos)
		for index, photo := range album.Photos {
			if photo == unescapePhotoName {
				targetPhoto = photoName
				if index != 0 {
					prevPhoto = album.Photos[index-1]
				}
				if index != photosLen-1 {
					nextPhoto = album.Photos[index+1]
				}
				break
			}
		}
		if targetPhoto == "" {
			return errors.New("photo not found")
		}

		htmlResponse, err := render.ParsePhoto(galleries, album.Name, targetPhoto, prevPhoto, nextPhoto)

		if err != nil {
			return err
		}

		return c.HTML(http.StatusOK, htmlResponse)
	})

	e.Static("/galleries", "galleries")

	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":1324"))
}
