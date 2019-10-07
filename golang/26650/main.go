package main

import (
	"os"
        "log"
)

func main() {
	f, err := os.Create("foo.txt")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer f.Close()

	_, _ = f.Write([]byte("foo-bar"))

	if err := f.Sync(); err != nil {
		log.Fatalf("Failed to invoke Fsync: %v", err)
	}
}
