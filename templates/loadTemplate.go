package main

import (
	"html/template"
	"log"
	"net/http"
)

// Arbitrary direction for the templates
const basePath = "templates"

// this is a server now
func main() {
	templates := loadHttpTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// basically, whenever the page will load
		// this trims the URL and looks for the file
		requestedFile := r.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/img", http.FileServer(http.Dir("public")))
	http.Handle("/css", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}

func loadHttpTemplates() *template.Template {
	result := template.New("templates")
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
