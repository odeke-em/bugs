package main

import (
	"log"
	"net/http"
	"time"
)

func waitItOut(w http.ResponseWriter, r *http.Request) {
	<-time.After(10 * time.Minute)
}

func main() {
	http.Handle("/", http.TimeoutHandler(http.HandlerFunc(waitItOut), time.Millisecond, `{"terminated": true}`))
	if err := http.ListenAndServe(":8899", nil); err != nil {
		log.Fatal(err)
	}
}
