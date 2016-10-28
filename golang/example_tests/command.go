package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("sh", "-c", "echo hey;echo 1>&2 stderr")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
