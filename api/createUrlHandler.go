package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/raiesbo/url-minifier/types"
)

type HandlerBody struct {
	URL string `json:"url"`
}

func HandleCreateNewURL(w http.ResponseWriter, r *http.Request) {
	var handlerBody HandlerBody

	err := json.NewDecoder(r.Body).Decode(&handlerBody)
	if err != nil {
		log.Fatalf("The content of the body is not correct: %s", err)
	}

	fmt.Println("handlerBody", handlerBody)

	// Validate URL
	_, err = url.ParseRequestURI(handlerBody.URL)
	if err != nil {
		log.Fatalf("The URL is not valid: %s", err)
	}

	// Create Short URL

	// Store in DB

	// Return URL

	var longURL types.URL
	_ = longURL

	fmt.Print("Hello")
}
