package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting on port 8080...")

	// Make a router. Aka Http request multiplexer
	mux := http.NewServeMux()

	// Define a handler to run if a request to the webserver
	mux.HandleFunc("/{$}", homeHandler)

	// This runs on the specified port with the router we defined above
	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The supermoto tutorial is now live!")
}
