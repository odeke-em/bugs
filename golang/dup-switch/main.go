package main

import (
	"io"
)

func whatis(x interface{}) string {
	switch x.(type) {
	case int:
		return "int"
	case int: // ERROR "duplicate"
		return "int8"
	case io.Reader:
		return "Reader1"
	case io.Reader: // ERROR "duplicate"
		return "Reader2"
	case interface {
		r()
		w()
	}:
		return "rw"
	case interface { // ERROR "duplicate"
		w()
		r()
	}:
		return "wr"

	}
	return ""
}

func main() {}
