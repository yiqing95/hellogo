package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		// post 请求
		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}
		// 用模型数据做一些事情
		_ = details
		result := struct {
			Success bool
			Msg     string
		}{true, fmt.Sprint(details)}
		// tmpl.Execute(w, struct{ Success bool }{true})
		tmpl.Execute(w, result)
	})
	http.ListenAndServe(":80", nil)
}
