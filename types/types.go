package types

import "time"

type URL struct {
	originalURL  string
	shortURL     string
	urlKey       string
	secretKey    string
	creationDate time.Time
}

func createURL(longURL string) URL {
	// Create URL Key

	// Create short URL

	// Create secret

	//Create URL object
	newURL := URL{
		originalURL:  longURL,
		creationDate: time.Now().UTC(),
	}
	return newURL
}
