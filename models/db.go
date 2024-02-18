package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/raiesbo/url-minifier/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBInstance struct {
	Uri        string
	Client     *mongo.Client
	Collection *mongo.Collection
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
	i.Collection = client.Database(utils.DATABASE).Collection(utils.COLLENTION)
}

func (i *DBInstance) StoreURL(url URL) {
	result, err := i.Collection.InsertOne(context.TODO(), url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}

func (i *DBInstance) FindURL(urlKey string) URL {
	var result URL
	err := i.Collection.FindOne(context.TODO(), bson.D{{"url_key", urlKey}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the urlKey %s\n", urlKey)
	}
	if err != nil {
		panic(err)
	}

	return result
}

func (i *DBInstance) UpdateURLCounter(url URL) {
	id, _ := primitive.ObjectIDFromHex(url.Id)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$inc", bson.D{{"clicks", 1}}}}

	// Updates the first document that has the specified "_id" value
	_, err := i.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}
