package pkg

import (
	"fmt"
	"net/http"
	"text/template"
)

func ServeTemplates(w http.ResponseWriter, data any, templatePaths []string) error {
	// If there is a base template, it should be first in templatePaths
	t, err := template.ParseFiles(templatePaths...)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		fmt.Println("Error parsing template:", err)
		return err
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		fmt.Println("Error executing template:", err)
		return err
	}
	return nil
}
