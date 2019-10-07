package main

import (
	"fmt"
	"sync"
)

func main() {
	mm := make(map[string]string)

	mm["0"] = "1"

	var wg = sync.WaitGroup{}
	for i := 0; i < 100000000; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(mm["0"], i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("done")
}
