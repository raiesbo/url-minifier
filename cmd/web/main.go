package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/raiesbo/servertools"
	"github.com/raiesbo/url-minifier/internal/db"
	"github.com/raiesbo/url-minifier/internal/url"
	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	redisClient := db.NewRedis(os.Getenv("DB_URI"))
	urlRepo := url.NewRedisRepository(redisClient)
	urlService := url.NewService(urlRepo)

	defer func(redisClient *redis.Client) {
		err := redisClient.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(redisClient)

	app := application{
		url: urlService,
		Tools: servertools.Tools{
			TmplsDir:    "./ui/html/pages/",
			BaseTmplDir: "./ui/html/partials/layout.tmpl",
		},
	}

	serverPort := os.Getenv("PORT")
	log.Printf("Listening to Port %v", serverPort)
	if err := http.ListenAndServe(":"+serverPort, rateLimiter(app.routes(), rate.Limit(2), 10)); err != nil {
		log.Fatal(err)
	}
}
