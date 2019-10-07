package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.Printf("PID: %d", os.Getpid())
	<-time.After(4 * time.Second)
}
