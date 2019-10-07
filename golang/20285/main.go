package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var cmd *exec.Cmd
	var err error

	cmd = exec.Command("/bin/ls")
	cmd.Dir = "/this/does/not/exist"
	err = cmd.Run()
	fmt.Printf("error was: %v\n", err)

	cmd = exec.Command("/bin/ls")
	cmd.Dir = "/this/does/not/exist"
	err = cmd.Run()
	fmt.Printf("error was: %v\n", err)
}
