package main

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func checkContention(b *testing.B) {
	n := runtime.NumCPU()
	if n < 2 {
		b.Skipf("Need at least 2 maxProcs, got %d", n)
	}
	var mu sync.Mutex

	shared := new(int)
	// The goal is to have n goroutines with
	// all contending for the same mutex.
	contend := func(w *sync.WaitGroup, period time.Duration, closeChan chan bool) {
		defer w.Done()

		for i := 0; ; i++ {
			select {
			case <-time.After(period):
				// Now contend
				mu.Lock()
				*shared = i
				mu.Unlock()
			case <-closeChan:
				return
			}
		}
	}

	syncPeriod := 20 * time.Nanosecond
	var wg sync.WaitGroup
	closeChan := make(chan bool)
	for i := 0; i < n-1; i++ {
		wg.Add(1)
		go contend(&wg, syncPeriod, closeChan)
	}

	// Now here on the main contend too
	for i := 0; i < n*40; i++ {
		<-time.After(syncPeriod)
		mu.Lock()
		*shared = i
		mu.Unlock()
	}
	close(closeChan)

	wg.Wait()
}

func BenchmarkContention(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkContention(b)
	}
}

func BenchmarkMutexLockUnlock(b *testing.B) {
	var mu sync.Mutex

	for i := 0; i < b.N; i++ {
		mu.Lock()
		mu.Unlock()
	}
}
