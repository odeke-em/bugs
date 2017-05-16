package main

import (
	"log"
	"net/http"

	"html/template"
)

var Handler http.Handler

func init() {
	Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, err := template.ParseFiles("x.tpl")
		if err != nil {
			log.Fatalln(err)
		}

		p.ExecuteTemplate(w, "x.tpl", nil)
	})
}