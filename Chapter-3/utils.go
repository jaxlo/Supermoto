package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func ServeTemplates(w http.ResponseWriter, data any, templatePaths []string) {
	// If there is a base template, it should be first in templatePaths
	t, err := template.ParseFiles(templatePaths...)
	if err != nil {
		http.Error(w, "Error parsing templates: "+err.Error(), http.StatusInternalServerError)
		fmt.Println("err parse")
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing templates: "+err.Error(), http.StatusInternalServerError)
		fmt.Println("err exec")
		return
	}
}
