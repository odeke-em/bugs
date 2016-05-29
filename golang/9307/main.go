package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/echo")
	pr, pw := io.Pipe()

	cmd.Stdin = pr
	if err := cmd.Start(); err != nil {
		log.Fatalf("Start: %v", err)
	}
	wrote := make(chan bool)
	go func() {
		defer close(wrote)
		fmt.Fprint(pw, "echo\n")
		_ = pw.Close()
	}()
	if err := cmd.Wait(); err != nil {
		log.Fatalf("Wait: %v", err)
	}
	<-wrote
}
