package main

import (
	"fmt"
	"net/http"
)

type eddie struct {
	chinesename string
}

func (e eddie) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm %s and I'm so hot...", e.chinesename)
}

func main() {
	e := eddie{
		"仔仔",
	}
	http.ListenAndServe(":8080", e)
}
