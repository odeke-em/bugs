package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	r := io.LimitReader(rand.Reader, 32)
	slurp, _ := ioutil.ReadAll(r)

	fmt.Printf("eq: %v\n", bytes.EqualFold(slurp, slurp))
}
