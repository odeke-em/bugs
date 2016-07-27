package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(strings.Repeat("a", 3*1024*1024*1024))
}
