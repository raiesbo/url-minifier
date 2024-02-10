package api

import (
	"fmt"
	"net/http"

	"github.com/raiesbo/url-minifier/types"
)

func HandleCreateNewURL(w http.ResponseWriter, r *http.Request) {
	var url types.URL
	_ = url

	fmt.Print("Hello")
}
