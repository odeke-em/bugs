package main

import "sync"

var sink int
var stopOnce sync.Once

func contend(i int) {
	stopOnce.Do(func() {
		sink = i
	})
}

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

	semaphore := make(chan bool, 100)
	for i := 1; i <= 1e5; i++ {
		for j := 0; j < 1e4; j++ {
			semaphore <- true

			wg.Add(1)
			go func(value int) {
				defer wg.Done()
				contend(value)
				<-semaphore
			}(j * i)
		}
	}
}
