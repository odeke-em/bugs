package main

import (
	"log"
	"os"
)

func main() {
	prc, pwc, err := os.Pipe()
	if err != nil {
		log.Fatalf("Failed to create os.Pipe(): %v", err)
	}
	bs := make([]byte, 10)
	if _, err := prc.Read(bs); err != nil {
		log.Fatalf("Failed to perform read: %v", err)
	}
	defer pwc.Close()
}
