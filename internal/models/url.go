package models

import (
	"time"

	"github.com/raiesbo/url-minifier/internal/utils"
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
	shortURL := host + "/" + urlKey

	// Create secret
	secretKey := utils.CreateURLKey(15)

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
