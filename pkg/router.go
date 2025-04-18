package pkg

import (
	"fmt"
	"net/http"
)

func RunRouter() {

	// Make a router. Aka Http request multiplexer
	mux := http.NewServeMux()

	// Make a handler for serving static files
	mux.HandleFunc("GET /static/", staticFiles)

	// Register what handler should be used for each matching URL
	mux.HandleFunc("/{$}", homeHandler)
	mux.HandleFunc("/section/getting_started", sectionGetting_started)
	mux.HandleFunc("/section/static", SectionStatic)
	// mux.HandleFunc("/section/apis", sectionApis)
	mux.HandleFunc("/section/templates", sectionTemplates)
	mux.HandleFunc("/section/forms", sectionForms)

	fmt.Println("Starting on port 8080...")
	http.ListenAndServe(":8080", mux)
}
