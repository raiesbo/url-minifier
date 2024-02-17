package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/raiesbo/url-minifier/models"
)

var DbInstance *models.DBInstance

func ScopeDBInstance(instance *models.DBInstance) {
	DbInstance = instance
}

type HandlerBody struct {
	URL string `json:"url"`
}

func HandleCreateNewURL(w http.ResponseWriter, r *http.Request) {
	var handlerBody HandlerBody

	err := json.NewDecoder(r.Body).Decode(&handlerBody)
	if err != nil {
		log.Fatalf("The content of the body is not correct: %s", err)
	}

	// Validate URL
	_, err = url.ParseRequestURI(handlerBody.URL)
	if err != nil {
		log.Fatalf("The URL is not valid: %s", err)
	}

	// Create Short URL
	newURL := models.NewURL(handlerBody.URL, r.Host)

	// Store in DB
	DbInstance.StoreURL(newURL)

	// Return URL
	urlJson, err := json.Marshal(newURL)
	if err != nil {
		log.Fatalf("Unable to transform to JSON: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(urlJson)
}
