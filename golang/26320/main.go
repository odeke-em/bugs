package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() { f1(); f2() }

func f1() {
	defer printLine("return")
	return
}

func f2() {
	defer printLine("panic")
	panic(1)
}

func printLine(s string) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("%10s: %s:%d\n", s, filepath.Base(file), line)
}
