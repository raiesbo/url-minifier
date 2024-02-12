package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleRedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Get urlKey
	urlKey := chi.URLParam(r, "urlKey")

	fmt.Println(urlKey)

	// Find URL from urlKey

	// Redirect
}
