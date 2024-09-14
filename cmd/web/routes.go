package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Get("/", HandleHome)
	mux.Get("/{urlKey}", HandleRedirectHandler)
	mux.Post("/api/url", HandleCreateNewURL)
	mux.Get("/api/admin", HandleCreateNewURL)
	mux.Delete("/api/admin", HandleCreateNewURL)

	// To serve static files from the correct folder
	fileServe := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/ui/static/*", http.StripPrefix("/ui/static", fileServe))

	return mux
}
