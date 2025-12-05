package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.Host, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

var ipLimiter sync.Map

func rateLimiter(next http.Handler, limit rate.Limit, burst int) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIpAddress(r)

		limiterAny, _ := ipLimiter.LoadOrStore(ip, rate.NewLimiter(limit, burst))
		limiter := limiterAny.(*rate.Limiter)
		if !limiter.Allow() {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			if err := json.NewEncoder(w).Encode(map[string]string{"error": "too many requests"}); err != nil {
				log.Fatal(err)
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getIpAddress(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return host
}
