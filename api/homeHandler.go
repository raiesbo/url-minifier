package api

import (
	"fmt"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wellcome to URL Minifier")
}
