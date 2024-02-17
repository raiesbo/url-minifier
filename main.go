package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/raiesbo/url-minifier/api"
	"github.com/raiesbo/url-minifier/config"
	"github.com/raiesbo/url-minifier/models"
)

var app config.AppConfig

var dbInstance models.DBInstance

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbInstance = models.NewDBInstance(os.Getenv("DB_URI"))
	dbInstance.Connect()
	api.ScopeDBInstance(&dbInstance)

	app.Env = os.Getenv("ENV")

	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Get("/", api.HandleHome)
	mux.Post("/url", api.HandleCreateNewURL)
	mux.Get("/{urlKey}", api.HandleRedirectHandler)
	mux.Get("/admin", api.HandleCreateNewURL)
	mux.Delete("/admin", api.HandleCreateNewURL)

	serverPort := os.Getenv("PORT")

	log.Printf("Listening to Port %v", serverPort)

	err := http.ListenAndServe(":"+serverPort, mux)
	log.Fatal(err)
}
