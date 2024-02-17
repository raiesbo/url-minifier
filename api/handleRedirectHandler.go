package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleRedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Get urlKey
	urlKey := chi.URLParam(r, "urlKey")

	// Find URL from urlKey
	result := DbInstance.FindURL(urlKey)

	// Redirect
	http.Redirect(w, r, result.OriginalURL, http.StatusSeeOther)
}
