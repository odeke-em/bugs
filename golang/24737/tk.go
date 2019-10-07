package main

import (
	"fmt"
	"net/http"
)

func main() {
	_, err := http.Get("https://golang.org/")
	fmt.Println(err)
}
