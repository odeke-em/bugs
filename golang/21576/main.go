package main

import _ "os/signal"

func main() {
	c := make(chan int)
	c <- 1
}
