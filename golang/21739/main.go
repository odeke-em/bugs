package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", `--profile-directory="Default"`)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("data=%s\n", out)
}
