package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "build", `-gcflags="-L -e"`, "-o=bar", "/foo")
	fmt.Println(cmd.String())
}
