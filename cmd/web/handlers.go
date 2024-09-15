package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/raiesbo/url-minifier/internal/models"
)

type HandlerBody struct {
	URL    string `json:"url"`
	Origin string `json:"origin"`
}

func (app *application) handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./ui/html/home.page.htm"))
	tmpl.Execute(w, nil)
}

func (app *application) handleCreateNewURL(w http.ResponseWriter, r *http.Request) {
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
	app.urls.StoreURL(newURL)

	// Send back HTML template
	if handlerBody.Origin == "HomePage" {
		html := fmt.Sprintf(`<div class="output">
		<p>%s</p>
		<p>%s</p>
		<p>%s</p>
		</div>`, newURL.OriginalURL, newURL.ShortURL, newURL.SecretKey)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(html))
		return
	}

	// Return URL
	urlJson, err := json.Marshal(newURL)
	if err != nil {
		log.Fatalf("Unable to transform to JSON: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(urlJson)
}
func (app *application) handleRedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Get urlKey
	urlKey := chi.URLParam(r, "urlKey")

	// Find URL from urlKey
	result, err := app.urls.FindURL(urlKey)
	if err != nil {
		panic(err)
	}

	// Update clics counter
	app.urls.UpdateURLCounter(result.Id)

	// Redirect
	http.Redirect(w, r, result.OriginalURL, http.StatusSeeOther)
}
