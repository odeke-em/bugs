package main

import (
	"encoding/asn1"
	"fmt"
	"time"
)

type A struct {
	t time.Time
}

func panics() {
	b, err := asn1.Marshal(A{})
	fmt.Printf("b: %s err: %v\n", b, err)
}

func noPanics() {
	b, err := asn1.Marshal(&A{})
	fmt.Printf("noPanics: b: %s err: %v\n", b, err)
}

func main() {
	noPanics()
	panics()
}
