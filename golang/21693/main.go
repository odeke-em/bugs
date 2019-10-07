package main

func main() {
	var s uint = 10
_ = 2.2 << s
	a := make([]byte, 1.0<<s) // valid
	_ = a[2.2<<s]             // accepted by go/types, not accepted by cmd/compile
}
