package main

import (
	"log"
	"net/http"
	"fmt"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	clientGone := w.(http.CloseNotifier).CloseNotify()
	w.Header().Set("Content-Type", "text/plain")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		fmt.Fprintf(w, "%v Hello\n", time.Now())
		w.(http.Flusher).Flush()
		select {
		case <-ticker.C:
		case <-clientGone:
			log.Printf("Client %v disconnected from the clock", r.RemoteAddr)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServeTLS(":1333", "key.crt", "key.key", nil); err != nil {
		log.Fatal(err)
	}
}
