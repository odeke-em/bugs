package main

func main() {
	var s uint
	var x = interface{}(1<<s + 1<<s)
	println(x)
}
