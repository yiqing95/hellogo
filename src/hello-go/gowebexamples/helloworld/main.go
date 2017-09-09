package main

import "net/http"

import "fmt"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello web\n"))

		fmt.Fprintf(w, "Hello you 'v requested : %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
