package repository

import (
	"context"
	"github.com/fwchen/saury/model"
	"github.com/juju/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewGalleryRepository(database *Database) *GalleryRepository {
	return &GalleryRepository{
		database: database,
	}
}

type GalleryRepository struct {
	database *Database
}

func (g *GalleryRepository) Save(album *model.Album) error {
	_, err := g.database.MongoClient.Collection("album").InsertOne(context.Background(), album)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (g *GalleryRepository) FindByName(name string) (*model.Album, error) {
	var album *model.Album
	err := g.database.MongoClient.Collection("album").FindOne(context.Background(), bson.M{
		"name": name,
	}).Decode(&album)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (g *GalleryRepository) FindAll(limit, offset int64) ([]model.Album, error) {
	cur, err := g.database.MongoClient.Collection("album").Find(context.Background(), bson.M{}, &options.FindOptions{
		Limit: func(i int64) *int64 { return &i }(limit),
		Skip:  func(i int64) *int64 { return &i }(offset),
	})
	defer cur.Close(context.Background())
	albums := make([]model.Album, 0)

	if err != nil {
		return nil, errors.Trace(err)
	}
	for cur.Next(context.Background()) {
		var album model.Album
		err := cur.Decode(&album)
		if err != nil {
			return nil, errors.Trace(err)
		}
		albums = append(albums, album)
	}
	return albums, nil
}

func (g *GalleryRepository) FindPhotos(albumName string, limit int, offset int) ([]model.Photo, error) {

	matchStage := bson.D{{"$match", bson.D{{"name", albumName}}}}
	unwindStage := bson.D{{"$unwind", "$photos"}}
	limitStage := bson.D{{"$limit", limit}}
	skipStage := bson.D{{"$skip", offset}}
	projectStage := bson.D{{"$project", bson.D{{"name", "$photos"}}}}

	cur, err := g.database.MongoClient.Collection("album").Aggregate(context.Background(), mongo.Pipeline{matchStage, unwindStage, limitStage, skipStage, projectStage})
	if err != nil {
		return nil, errors.Trace(err)
	}

	photos := make([]model.Photo, 0)
	err = cur.All(context.Background(), &photos)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return photos, nil
}

func (g *GalleryRepository) FindPhotosCount(albumName string) (int, error) {
	var album *model.Album
	err := g.database.MongoClient.Collection("album").FindOne(context.Background(), bson.M{}).Decode(&album)
	if err != nil {
		return 0, errors.Trace(err)
	}
	pageCount := len(album.Photos)
	return pageCount, nil
}
