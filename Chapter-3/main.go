package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting webserver on port 8000")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /static/", StaticFiles)
	mux.HandleFunc("GET /{$}", Home)

	http.ListenAndServe(":8000", mux)
}
