package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

type eddie string

func (e eddie) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		r.Method,
		r.Form,
	}
	tpl.ExecuteTemplate(w, "go.html", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("go.html"))
}

func main() {
	var e eddie
	http.ListenAndServe(":8080", e)
}
