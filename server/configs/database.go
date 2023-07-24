package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database = ConnectDB()

func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(Env.MONGODB_URI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func DisconnectDB() error {
	return database.Disconnect(context.Background())
}

func GetCollection(collectionName string) *mongo.Collection {
	return database.Database("stas").Collection(collectionName)
}
