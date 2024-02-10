package main

import (
	"flag"
	"net/http"

	"github.com/raiesbo/url-minifier/api"
)

func main() {
	listenAddr := flag.String("listenaddr", ":49999", "todo")
	flag.Parse()

	http.HandleFunc("/url", api.HandleCreateNewURL)
	http.HandleFunc("/user", api.HandleCreateNewURL)
	http.HandleFunc("/user", api.HandleCreateNewURL)
	http.HandleFunc("/user", api.HandleCreateNewURL)

	http.ListenAndServe(*listenAddr, nil)
}
