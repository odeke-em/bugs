package main

import (
	"log"
	"os/exec"
)

func main() {
	output, err := exec.Command("go", "version").CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run os/exec: %v", err)
	}
	log.Printf("%s\n", output)
}
