package main

import (
	"net/http"
)

func StaticFiles(w http.ResponseWriter, r *http.Request) {
	// Create a file server for the specified directory
	fs := http.FileServer(http.Dir("staticfiles"))
	// And serve those files
	http.StripPrefix("/static", fs).ServeHTTP(w, r)
}

func Home(w http.ResponseWriter, r *http.Request) {
	// Optionally put the []string{} in it's own var
	ServeTemplates(w, nil, []string{"views/base.html", "views/home.html"})
}
