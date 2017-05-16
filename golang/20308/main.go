package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("./sample", "pick", "up", "the", "phone")
	cmd.Dir = "."
	outBytes, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("outBytes: %s\n", outBytes)
}
