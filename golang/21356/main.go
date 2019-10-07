package main

import (
	"math/rand"
	"sync"
	"time"
)

func spawnThem(wg sync.WaitGroup, period time.Duration) {
	defer wg.Done()

	sp := time.NewTicker(period)
	for {
		<-sp.C
	}
}

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		spawnThem(wg, time.Duration(rand.Intn(30))*time.Second)
	}
}
