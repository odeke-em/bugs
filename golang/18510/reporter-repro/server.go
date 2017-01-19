package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, _ := strconv.Atoi(r.URL.Query().Get("num"))
		log.Println("new request:", n)
		w.Header().Set("Content-Type", "text/plain")
		s := strings.Repeat("hello,world", n) //
		w.Write([]byte(s))
		w.(http.Flusher).Flush()
		clientGone := w.(http.CloseNotifier).CloseNotify()
		<-clientGone
		log.Printf("Client %v closed", r.RemoteAddr)
	})
	if err := http.ListenAndServeTLS(":7001", "cert.pem", "privkey.pem", nil); err != nil {
		log.Fatal(err)
	}
}
