package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// tmpl := template.Must(template.ParseFiles("base.tmpl", "home.tmpl"))
	tmpl := template.Must(template.ParseFiles("views/layouts/base.html", "views/home.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if err := tmpl.Execute(w, nil); err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":80", nil)
}
