package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	srv := http.Server{
		Addr: ":8001",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/index" {
				w.Write([]byte("OK"))
			} else {
				w.Header().Add("Location", "http://servername/index")
				w.WriteHeader(http.StatusFound)
			}
		}),
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	log.Fatal(http.ListenAndServe(":8000", &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = "http"
			r.URL.Host = "127.0.0.1:8001"
		},
		ModifyResponse: func(res *http.Response) error {
log.Printf("done here\n")
			res.Header.Set("Location", "/info")
			return nil
		},
	}))
}
