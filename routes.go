package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/raiesbo/url-minifier/api"
	"github.com/raiesbo/url-minifier/pages"
)

func routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Get("/", pages.HandleHome)
	mux.Get("/{urlKey}", api.HandleRedirectHandler)
	mux.Post("/api/url", api.HandleCreateNewURL)
	mux.Get("/api/admin", api.HandleCreateNewURL)
	mux.Delete("/api/admin", api.HandleCreateNewURL)

	// To serve static files from the correct folder
	fileServe := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServe))

	return mux
}
