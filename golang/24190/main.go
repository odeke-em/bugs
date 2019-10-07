package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://gopkg.in/sourcemap.v1?go-get=1")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", res)
}