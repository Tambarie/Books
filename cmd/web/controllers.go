package main

import (
	"html/template"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	templateFiles := []string{
		"../../ui/html/home.page.html",
		"../../ui/html/base.layout.html",
		"../../ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
