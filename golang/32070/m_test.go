package main

import (
	"fmt"
	"runtime"
	"testing"
)

func main() {
	fmt.Println(runtime.Version())
}

func f(s string) error {
	return fmt.Errorf("abc %w", s)
}

func TestIt(t *testing.T) {
}
