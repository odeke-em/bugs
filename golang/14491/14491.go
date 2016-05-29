package main

import (
	"bytes"
	"io/ioutil"
	"log"
)

func main() {
	outfilePath := "temp"
	content := []byte("foo")
	err := ioutil.WriteFile(outfilePath, content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	retrieved, err := ioutil.ReadFile(outfilePath)
	if err != nil {
		log.Fatal(err)
	}

	if !bytes.Equal(retrieved, content) {
		log.Fatalf("file content doesn't match, read %s expected %s\n", retrieved, content)
	}
}
