package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("error"))
			return
		}
		params := req.Form
		fmt.Println("params:", params, "count:", len(params))
		for key, values := range params {
			fmt.Println("param", key, ":", values)
		}
		w.Write([]byte("OK"))
	})

	fmt.Println("starting on port 9999")
	server := &http.Server{
		Addr: ":9999",
	}
	server.ListenAndServe()
}