package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	//create a client connecting to the database by the .env URI
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GetEnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	//Ping database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

var DB *mongo.Client = ConnectDB()

// Get a collection from the database
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("virtual-pets").Collection(collectionName)
	if collection == nil {
		log.Fatal("No Collection found")
	}
	return collection
}
