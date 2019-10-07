package main

func f1() func() {
	return f2
}

func f2() {
	println("hey")
}

func main() {
    f1()
}
