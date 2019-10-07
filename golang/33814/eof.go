package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	defer fmt.Printf("Spent: %s\n", time.Since(time.Now()))

	var b [1]byte
	_, err := os.Stdin.Read(b[:])
	if err != nil {
		log.Fatalf("Failed to read a byte from stdin: %v", err)
	}
}
