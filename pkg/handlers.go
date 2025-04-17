package pkg

import "net/http"

func staticFiles(w http.ResponseWriter, r *http.Request) {
	// Create a file server for the specified directory
	fs := http.FileServer(http.Dir("staticfiles"))
	// And serve those files
	http.StripPrefix("/static", fs).ServeHTTP(w, r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	ServeTemplates(w, nil, []string{"templates/base.html", "templates/home.html"})
}

func sectionGetting_started(w http.ResponseWriter, r *http.Request) {
	ServeTemplates(w, nil, []string{"templates/base.html", "templates/sections/getting_started.html"})
}

func SectionStatic(w http.ResponseWriter, r *http.Request) {
	ServeTemplates(w, nil, []string{"templates/base.html", "templates/sections/static.html"})
}

func sectionApis(w http.ResponseWriter, r *http.Request) {
	ServeTemplates(w, nil, []string{"templates/base.html", "templates/sections/apis.html"})
}

func sectionTemplates(w http.ResponseWriter, r *http.Request) {
	ServeTemplates(w, nil, []string{"templates/base.html", "templates/sections/templates.html"})
}

func sectionForms(w http.ResponseWriter, r *http.Request) {
	ServeTemplates(w, nil, []string{"templates/base.html", "templates/sections/forms.html"})
}
