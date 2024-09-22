package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	toolkit "github.com/raiesbo/server-toolkit"
	"github.com/raiesbo/url-minifier/internal/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	app := application{
		urls:  &models.UrlModel{},
		users: &models.UserModel{},
		tk: &toolkit.Toolkit{
			TmplsDir:    "./ui/html/pages/",
			BaseTmplDir: "./ui/html/partials/layout.tmpl",
		},
	}

	app.connect(os.Getenv("DB_URI"))

	serverPort := os.Getenv("PORT")
	log.Printf("Listening to Port %v", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, app.routes()))
}
