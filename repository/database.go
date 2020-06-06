package repository

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	MongoClient *mongo.Database
}

func Connect(mongoUri string, mongoDB string) (*Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}
	return &Database{
		MongoClient: client.Database(mongoDB),
	}, nil
}

