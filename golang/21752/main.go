package main

var A = func() int {
	if false {
		return A
	}
	return 0
}()

func main() {
}
