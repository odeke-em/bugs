package main

import (
	"archive/tar"
	"os"
)

func main() {
	input, err := os.Open("extended_header.tar")
	if err != nil {
		panic(err)
	}

	tarR := tar.NewReader(input)
	_, err = tarR.Next()
	panic(err) // should be nil, but is "invalid header"
}