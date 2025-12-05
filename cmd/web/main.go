package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/raiesbo/servertools"
	"github.com/raiesbo/url-minifier/internal/models"
	"golang.org/x/time/rate"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	app := application{
		urls:  &models.UrlRepository{},
		users: &models.UserModel{},
		Tools: servertools.Tools{
			TmplsDir:    "./ui/html/pages/",
			BaseTmplDir: "./ui/html/partials/layout.tmpl",
		},
	}

	app.connect(os.Getenv("DB_URI"))

	serverPort := os.Getenv("PORT")
	log.Printf("Listening to Port %v", serverPort)
	if err := http.ListenAndServe(":"+serverPort, rateLimiter(app.routes(), rate.Limit(2), 10)); err != nil {
		log.Fatal(err)
	}
}
