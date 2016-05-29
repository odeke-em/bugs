package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/echo")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("StdinPipe: %v", err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatalf("Start: %v", err)
	}
	wrote := make(chan bool)
	go func() {
		defer close(wrote)
		fmt.Fprint(stdin, "echo\n")
	}()
	if err := cmd.Wait(); err != nil {
		log.Fatalf("Wait: %v", err)
	}
	<-wrote
}
