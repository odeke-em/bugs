package main

import (
	"fmt"
	"net/http"
)

func main() {
	try("ABCD\000") // nothing special
	try("OggS\000")	// legal beginning of application/ogg file
	try("owow\000")	// NOT an ogg file
	try("oooS\000")	// NOT an ogg file
}

func try(s string) {
	b := []byte(s)
	fmt.Println(s, ":", http.DetectContentType(b))
}
