package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	backendServer := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "this call was relayed by the reverse proxy")
		}),
	}
	fmt.Printf("starting app on port %s\n", port)
	err := backendServer.ListenAndServe()
	log.Fatal(fmt.Sprintf("BACKEND ERROR: %s", err))
}
