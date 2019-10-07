package main

import (
	"log"
	"reflect"
)

func main() {
	deref := reflect.ValueOf(new([0]byte)).Elem().Interface()
	var want [0]byte
	if g, w := reflect.DeepEqual(deref, want), true; g != w {
		log.Fatalf("got=%t want=%t", g, w)
	}
}
