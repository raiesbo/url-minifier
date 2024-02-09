package main

import (
	"flag"
	"net/http"

	"github.com/raiesbo/url-minifier/api"
)

func main() {
	listenAddr := flag.String("listenaddr", ":49999", "todo")
	flag.Parse()

	http.HandleFunc("/user", api.HandleGetUser)

	http.ListenAndServe(*listenAddr, nil)
}
