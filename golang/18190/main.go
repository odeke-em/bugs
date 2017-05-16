package main

import (
	"fmt"
	"net/http"
	"plugin"
)

func main() {
	p, err := plugin.Open("pg.so")

	if err != nil {
		panic(err)
	}

	sym, err := p.Lookup("Handler")

	if err != nil {
		panic(err)
	}

	fmt.Println("starting...")

	http.ListenAndServe(":8080", *(sym.(*http.Handler)))
}