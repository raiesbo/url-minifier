package main

import (
	"html/template"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./ui/html/home.page.htm"))
	tmpl.Execute(w, nil)
}
