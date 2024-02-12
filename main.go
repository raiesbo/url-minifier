package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/raiesbo/url-minifier/api"
)

const serverPort = ":8080"

func main() {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Get("/", api.HandleHome)
	mux.Post("/url", api.HandleCreateNewURL)
	mux.Get("/{urlKey}", api.HandleRedirectHandler)
	mux.Get("/admin", api.HandleCreateNewURL)
	mux.Delete("/admin", api.HandleCreateNewURL)

	log.Printf("Listening to Port %v", serverPort)

	err := http.ListenAndServe(serverPort, mux)
	log.Fatal(err)
}
