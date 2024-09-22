package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// To serve static files from the correct folder
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Routes
	mux.Handle("GET /{$}", logger(http.HandlerFunc(app.handleHome)))
	mux.Handle("POST /{$}", logger(http.HandlerFunc(app.handleCreateNewURL)))
	mux.Handle("GET /{urlKey}", logger(http.HandlerFunc(app.handleRedirectHandler)))

	return mux
}
