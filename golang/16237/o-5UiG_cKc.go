package main

import "bytes"
import "fmt"
import "strings"

func main() {
	N := int((^uint(0))/255 + 1)
	b := make([]byte, 255)
	fmt.Println(len(bytes.Repeat(b, N)))
	s := string(b)
	fmt.Println(len(strings.Repeat(s, N)))
}
