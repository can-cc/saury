package repository

import (
	"context"
	"github.com/fwchen/saury/model"
	"go.mongodb.org/mongo-driver/bson"
)

func NewGalleryRepository(database *Database) *GalleryRepository {
	return &GalleryRepository{
		database: database,
	}
}

type GalleryRepository struct {
	database *Database
}

func (g *GalleryRepository) save(album *model.Ａlbum) {
	g.database.MongoClient.Collection("album").InsertOne(context.Background(), album)
}

func (g *GalleryRepository) findByName(name string) (*model.Ａlbum, error) {
	var album *model.Ａlbum
	err := g.database.MongoClient.Collection("gallery").FindOne(context.Background(), bson.M{
		"name": name,
	}).Decode(album)
	if err != nil {
		return nil, err
	}
	return album, nil
}