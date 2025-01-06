package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func DBinstance() *mongo.Client {
	// MongoDb := "mongodb+srv://souvik741156:<db_password>@cluster0.l0bgq.mongodb.net/"
	MongoDb := "mongodb://localhost:27017"
	fmt.Println(MongoDb)

	client, err := mongo.Connect(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to database")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionname string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionname)
	return collection
}
