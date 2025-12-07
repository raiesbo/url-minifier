package main

import (
	"log"
	"net/http"
	"net/url"

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

	if err := app.RenderTmpl(w, r, "home.tmpl", http.StatusOK, data); err != nil {
		panic(err)
	}
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
	existingURL, err := app.url.FindByLongURL(form.LongURL)
	if err == nil {
		data := urlShortenForm{
			LongURL:  existingURL.Original,
			ShortURL: existingURL.Short,
		}
		if err = app.RenderTmpl(w, r, "home.tmpl", http.StatusAccepted, data); err != nil {
			log.Fatal(err)
		}
		return
	}

	newURL, err := app.url.NewURL(form.LongURL, r.Host)
	if err != nil {
		log.Fatalf("The URL is not valid: %s", err)
	}

	data := urlShortenForm{
		LongURL:  newURL.Original,
		ShortURL: newURL.Short,
	}

	if err = app.RenderTmpl(w, r, "home.tmpl", http.StatusAccepted, data); err != nil {
		log.Fatal(err)
	}
}

func (app *application) handleRedirectHandler(w http.ResponseWriter, r *http.Request) {
	urlKey := r.PathValue("urlKey")

	foundUrl, err := app.url.FindBySlug(urlKey)
	if err != nil {
		panic(err)
	}

	// Update redirects counter
	if err = app.url.UpdateCounter(foundUrl.Slug); err != nil {
		panic(err)
	}

	// Redirect
	http.Redirect(w, r, foundUrl.Original, http.StatusSeeOther)
}
