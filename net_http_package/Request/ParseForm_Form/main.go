package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("goto.html"))
}

type eddie int

func (e eddie) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	//tpl.Execute(w, req.Form)
	tpl.Execute(w, req.PostForm)
	//tpl.ExecuteTemplate(w, "goto.html"", req.Form)

	//req.Form will get value including both the URL
	// field's query parameters and the POST or PUT form data.

	//req.PostForm only get value from POST, PATCH,
	// or PUT body parameters.
}

func main() {
	var e eddie
	http.ListenAndServe(":8080", e)
}
