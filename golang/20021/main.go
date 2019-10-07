package main

const N = 1000 * 1000 * 537
var a [N]byte

func main() {
	// this will stack overflow
	for _, v := range a {
	  _ = v
	}

	// edit: blow solo also stack overflow
	var x [N]byte
	var y [N]byte
	_, _ = x, y
}