package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type URL struct {
	OriginalURL string    `bson:"original_url,omitempty"`
	ShortURL    string    `bson:"short_url,omitempty"`
	UrlKey      string    `bson:"url_key,omitempty"`
	SecretKey   string    `bson:"secret_key,omitempty"`
	Clicks      int       `bson:"clicks,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
}

func NewURL(longURL string, host string) URL {
	text := longURL

	hasher := md5.New()
	hasher.Write([]byte(text))

	urlKey := hex.EncodeToString(hasher.Sum(nil))[0:10]

	// Create short URL
	shortURL := host + "/" + urlKey

	// Create secret

	//Create URL object
	newURL := URL{
		OriginalURL: longURL,
		UrlKey:      urlKey,
		ShortURL:    shortURL,
		CreatedAt:   time.Now().Local(),
	}

	return newURL
}
