package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database

func ConnectMongoDB(dbUri, dbName string) error {
	// initial 10 second counter to expired context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// cancel after function call in order to prevent memory leak
	defer cancel()
	// connect to database
	clientOptions := options.Client().ApplyURI(dbUri)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	DB = client.Database(dbName)
	return nil
}
