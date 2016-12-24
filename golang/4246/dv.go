package main

import "fmt"

func main() {
	x := 0
	done := make(chan bool)
	go func() {
		x = 42
		done <- true
	}()
	x = 42
	<-done
	fmt.Println("Hello, playground", x)
}
