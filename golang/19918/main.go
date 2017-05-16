package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func execIt() error {
	cmd := exec.Command("some command")
	cmd.Stderr = ioutil.Discard
	cmd.Stdout = ioutil.Discard
	if _, err := cmd.StdoutPipe(); err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		return err
	}
	return cmd.Wait()
}

func main() {
	fmt.Printf("err: %v\n", execIt())
}
