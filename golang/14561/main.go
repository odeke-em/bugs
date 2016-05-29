package main

import (
	"log"
	"time"
)

func main() {
	log.Printf("Timer started")
	resultsChan := make(chan string)
	go func() {
		resultsChan <- "C for channel"
		close(resultsChan)
	}()

	select {
	case <-time.After(time.Second * 10):
		log.Printf("timer expired")
	case msg := <-resultsChan:
		log.Printf("msg here %s", msg)
	}
	log.Printf("Bye!")
}
