package main

import (
	"fmt"
	"runtime"
	"testing"
)

func main() {
	fmt.Println(runtime.Version())
}

func TestFoo(t *testing.T) {
	t.Fail("what?")
}