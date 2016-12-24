package main

import (
	"fmt"
	"os"
	"runtime"
)

//go:nowritebarrierrec
func hey() {
}

func main() {
	f, _ := os.Open("main.go")
	defer f.Close()

	fmt.Printf("%s\nname: %s", runtime.GOROOT(), f.Name())
}
