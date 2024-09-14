package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleRedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Get urlKey
	urlKey := chi.URLParam(r, "urlKey")

	// Find URL from urlKey
	result := DbInstance.FindURL(urlKey)

	// Update clics counter
	DbInstance.UpdateURLCounter(result)

	// Redirect
	http.Redirect(w, r, result.OriginalURL, http.StatusSeeOther)
}
