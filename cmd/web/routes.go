package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Get("/", app.handleHome)
	mux.Get("/{urlKey}", app.handleRedirectHandler)
	mux.Post("/api/url", app.handleCreateNewURL)
	mux.Get("/api/admin", app.handleCreateNewURL)
	mux.Delete("/api/admin", app.handleCreateNewURL)

	// To serve static files from the correct folder
	fileServe := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/ui/static/*", http.StripPrefix("/ui/static", fileServe))

	return mux
}
