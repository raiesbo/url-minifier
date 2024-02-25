package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/raiesbo/url-minifier/api"
	"github.com/raiesbo/url-minifier/models"
)

var dbInstance models.DBInstance

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	dbInstance = models.NewDBInstance(os.Getenv("DB_URI"))
	dbInstance.Connect()
	api.ScopeDBInstance(&dbInstance)

	serverPort := os.Getenv("PORT")
	log.Printf("Listening to Port %v", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, routes()))
}
