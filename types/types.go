package types

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type URL struct {
	OriginalURL string
	ShortURL    string
	UrlKey      string
	SecretKey   string
	Slicks      int
	CreatedAt   time.Time
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
