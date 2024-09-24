package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello shell!")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", root)
	http.ListenAndServe(":4000", mux)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello web!")
}
