package models

import (
	"context"
	"fmt"
	"time"

	"github.com/raiesbo/url-minifier/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type URL struct {
	Id          string    `bson:"_id,omitempty"`
	OriginalURL string    `bson:"original_url,omitempty"`
	ShortURL    string    `bson:"short_url,omitempty"`
	UrlKey      string    `bson:"url_key,omitempty"`
	SecretKey   string    `bson:"secret_key,omitempty"`
	Clicks      int       `bson:"clicks,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
}

func NewURL(longURL string, host string) URL {
	urlKey := utils.CreateURLKey(10)

	// Create short URL
	shortURL := fmt.Sprintf("%s/%s", host, urlKey)

	// Create secret
	secretKey := utils.CreateURLKey(10)

	//Create URL object
	newURL := URL{
		OriginalURL: longURL,
		UrlKey:      urlKey,
		ShortURL:    shortURL,
		SecretKey:   secretKey,
		Clicks:      0,
		CreatedAt:   time.Now().Local(),
	}

	return newURL
}

type UrlModel struct {
	DB *mongo.Collection
}

func (m *UrlModel) StoreURL(url URL) error {
	_, err := m.DB.InsertOne(context.TODO(), url)
	if err != nil {
		return err
	}
	return nil
}

func (m *UrlModel) FindByKey(urlKey string) (URL, error) {
	var result URL
	err := m.DB.FindOne(context.TODO(), bson.D{{Key: "url_key", Value: urlKey}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the urlKey %s\n", urlKey)
	}
	if err != nil {
		panic(err)
	}

	return result, nil
}

func (m *UrlModel) FindByLongURL(longURL string) (*URL, error) {
	var result *URL
	err := m.DB.FindOne(context.TODO(), bson.D{{Key: "original_url", Value: longURL}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the urlKey %s\n", longURL)
	// }
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *UrlModel) UpdateURLCounter(urlId string) error {
	id, _ := primitive.ObjectIDFromHex(urlId)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "clicks", Value: 1}}}}

	// Updates the first document that has the specified "_id" value
	_, err := m.DB.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	return nil
}
