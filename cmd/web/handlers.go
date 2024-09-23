package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/raiesbo/url-minifier/internal/models"
	"github.com/raiesbo/url-minifier/internal/validator"
)

type urlShortenForm struct {
	LongURL  string
	ShortURL string
	validator.Validator
}

func (app *application) handleHome(w http.ResponseWriter, r *http.Request) {
	data := urlShortenForm{
		LongURL: "",
	}

	app.RenderTmpl(w, r, "home.tmpl", http.StatusOK, data)
}

func (app *application) handleCreateNewURL(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	form := urlShortenForm{
		LongURL: r.FormValue("long-url"),
	}

	// Validate URL
	_, err := url.ParseRequestURI(form.LongURL)
	if err != nil {
		log.Fatalf("The URL is not valid: %s", err)
	}

	// Check if already exists and return
	existingURL, err := app.urls.FindByLongURL(form.LongURL)
	if err == nil {
		data := urlShortenForm{
			LongURL:  existingURL.OriginalURL,
			ShortURL: existingURL.ShortURL,
		}
		if err = app.RenderTmpl(w, r, "home.tmpl", http.StatusAccepted, data); err != nil {
			log.Fatal(err)
		}
		return
	}

	// Create Short URL
	newURL := models.NewURL(form.LongURL, r.Host)

	// Store in DB
	app.urls.StoreURL(newURL)

	data := urlShortenForm{
		LongURL:  newURL.OriginalURL,
		ShortURL: newURL.ShortURL,
	}

	if err = app.RenderTmpl(w, r, "home.tmpl", http.StatusAccepted, data); err != nil {
		log.Fatal(err)
	}
}
func (app *application) handleRedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Get urlKey
	urlKey := r.PathValue("urlKey")

	// Find URL from urlKey
	result, err := app.urls.FindByKey(urlKey)
	if err != nil {
		panic(err)
	}

	// Update clics counter
	app.urls.UpdateURLCounter(result.Id)

	// Redirect
	http.Redirect(w, r, result.OriginalURL, http.StatusSeeOther)
}
