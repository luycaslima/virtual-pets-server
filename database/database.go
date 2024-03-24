package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/luycaslima/virtual-pets-server/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	GetDb() *mongo.Client
}

// Connect to a Database
func ConnectDB() *mongo.Client {
	configs.LoadEnvFile()
	fmt.Println("Connecting Database")
	ctx := context.Background()
	//Because of the free host, timeout happens before it can activate itself
	//context.WithTimeout(context.TODO(), 2*time.Second)
	//defer cancel()

	//create a client connecting to the database from the .env URI
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		log.Fatal(err)
	}

	//Ping database
	//Pings make the resilience of the server go down
	//if the server is starting up, ping will give an error
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetCollectionFromDB(db *mongo.Client, collectionName string) *mongo.Collection {
	collection := db.Database(os.Getenv("DATABASE_NAME")).Collection(collectionName)
	if collection == nil {
		log.Fatal("No Collection Found!")
	}
	return collection
}
