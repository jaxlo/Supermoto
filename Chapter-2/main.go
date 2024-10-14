package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello shell!")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", root)
	mux.HandleFunc("GET /html", getHTML)
	mux.HandleFunc("GET /static/", getStatic)
	http.ListenAndServe(":8000", mux)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello web!")
}

func getHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/static.html")
}

func getStatic(w http.ResponseWriter, r *http.Request) {
	// Create a file server for the specified directory
	fs := http.FileServer(http.Dir("staticfiles"))
	// And serve those files
	http.StripPrefix("/static", fs).ServeHTTP(w, r)
}
