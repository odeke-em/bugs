package main

import (
	"errors"
	"fmt"
)

func serve(ich <-chan int) error {
	for {
		i := <-ich
		if i%2 == 3 {
			return errors.New("encountered here")
		}
		fmt.Printf("still running\n")
	}
}

func main() {
	ich := make(chan int, 1)
	go func() {
		defer fmt.Printf("done here!\n")
		defer close(ich)
		for i := 0; i < 10; i++ {
			ich <- i
		}
	}()

	err := serve(ich)
	fmt.Printf("err: %v\n", err)
}
