package main

import "time"

var a1 int
var b1 bool
var s1 string

func T1() {
	var x int
	var a int
	var s string

	if time.Now().Unix()%2 == 0 {
		x = 10
		s = "even"
		a = 42
	} else {
		a = 99
		s = "odd"
		x = 747
	}

	println(s)
	println(a)
	println(x)
}
