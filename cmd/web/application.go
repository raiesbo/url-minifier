package main

import (
	"context"
	"log"
	"time"

	"github.com/raiesbo/url-minifier/internal/models"
	"github.com/raiesbo/url-minifier/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	urls  *models.UrlModel
	users *models.UserModel
}

func (app *application) connect(uri string) {
	if uri == "" {
		log.Fatal("You must set your 'DB_URI' environment variable.")
	}

	opts := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	database := client.Database(utils.DATABASE)
	app.urls.DB = database.Collection(utils.COLLECTION_URLS)
	app.users.DB = database.Collection(utils.COLLECTION_USERS)
}
