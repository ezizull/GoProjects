package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstence exported
type MongoInstence struct {
	Client	*mongo.Client
	Db     	*mongo.Database
}

// Mongo exported
var Mongo MongoInstence

const (
	dbName 		= "goprojects"
	mongoURL 	= "mongodb://localhost:27017/" + dbName
)

// MongoConnect exported
func MongoConnect() error{
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil { return err }
	Mongo = MongoInstence {
		Client: client,
		Db:     db,
	}

	return nil
}