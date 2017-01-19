package main

func foo(prefix string, rest ...string) {
}

func main() {
	args := []string{"foo", "bar"}
	foo("ls", "info", args...)
}
