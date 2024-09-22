package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// To serve static files from the correct folder
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.Handle("GET /{$}", logger(http.HandlerFunc(app.handleHome)))
	mux.Handle("POST /{$}", logger(http.HandlerFunc(app.handleCreateNewURL)))
	// mux.Handle("GET /{urlKey}", logger(http.HandlerFunc(app.handleRedirectHandler)))
	// mux.Handle("GET /api/admin", logger(http.HandlerFunc(app.handleCreateNewURL)))
	// mux.Handle("DELETE /api/admin", logger(http.HandlerFunc(app.handleCreateNewURL)))

	return mux
}
