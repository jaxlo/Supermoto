package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting webserver...")

	// Make a router. Aka Http request multiplexer
	mux := http.NewServeMux()

	// Define a handler to run if a request to the webserver
	mux.HandleFunc("/{$}", homeHandler)
	mux.HandleFunc("/about", aboutHandler)

	// This runs on the specified port with the router we defined above
	http.ListenAndServe(":8000", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About the website")
}
