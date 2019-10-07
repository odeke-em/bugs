package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{Addr: ":8888"}
	srv.Close()

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Serving: %v", err)
	}
}
