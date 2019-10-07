package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
	d
)

const vf int = 0x3

func calc() int {
    if 1&3 != 1 {
        panic("Panicking here on 1")
    }
    if 2&0x3ff1 != 0 {
        panic("Panicking here on 2")
    }
    return 3
}

func main() {
	var v int
	v = a
	switch v {
	case a, b:
		fmt.Println("ab")
	case c, b:
		fmt.Println("bc")
	case d:
		fmt.Println("d")
	case 3:
		fmt.Println("3")
	case vf:
		fmt.Println("3")
	default:
		fmt.Println("other")
	}
}
