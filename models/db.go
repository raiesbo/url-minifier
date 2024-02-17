package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBInstance struct {
	Uri    string
	Client *mongo.Client
}

func NewDBInstance(dbURI string) DBInstance {
	return DBInstance{Uri: dbURI}
}

func (i *DBInstance) Connect() {
	if i.Uri == "" {
		log.Fatal("You must set your 'DB_URI' environment variable.")
	}

	opts := options.Client().ApplyURI(i.Uri)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	i.Client = client
}

func (i *DBInstance) StoreURL(url URL) {
	coll := i.Client.Database("URLMinifier").Collection("urls")

	result, err := coll.InsertOne(context.TODO(), url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}

func (i *DBInstance) FindURL(urlKey string) URL {
	coll := i.Client.Database("URLMinifier").Collection("urls")

	var result URL
	err := coll.FindOne(context.TODO(), bson.D{{"url_key", urlKey}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the urlKey %s\n", urlKey)
	}
	if err != nil {
		panic(err)
	}

	return result
}
