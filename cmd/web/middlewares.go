package main

import (
	"log"
	"net/http"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.Host, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
