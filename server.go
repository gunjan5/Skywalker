package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title string
}

var templates = template.Must(template.ParseFiles("index.html"))

func RootHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}
	templates.ExecuteTemplate(response, "index.html", Page{Title: "Home"})
}
