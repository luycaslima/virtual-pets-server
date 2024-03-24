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

type mongoDbDatabase struct {
	Db *mongo.Client
}

func ConnectMongoDB() *mongoDbDatabase {

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
	return &mongoDbDatabase{Db: client}
}

func (db *mongoDbDatabase) GetDb() *mongo.Client {
	return db.Db
}
