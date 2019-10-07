package main

import (
	"log"
	"os"
	"runtime/debug"
	"runtime/trace"
	"sync"
)

func main() {
	f, err := os.Create("pt.txt")
	if err != nil {
		log.Fatalf("Failed to create trace profile file: %v", err)
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	n := 8
	var wg sync.WaitGroup
	defer wg.Wait()

	var rl, wl []*os.File
	for i := 0; i < n; i++ {
		r, w, err := os.Pipe()
		if err != nil {
			for j := 0; j < i; j++ {
				_ = rl[j].Close()
				_ = wl[j].Close()
			}
			log.Fatalf("Invoking os.Pipe() #%d: %v", i, err)
		}
		rl = append(rl, r)
		wl = append(wl, w)
	}

	if false {
		defer debug.SetMaxThreads(debug.SetMaxThreads(8))
	}

	reading := make(chan bool, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ri := rl[i]
			reading <- true
			bs := make([]byte, 1)
			if _, err := ri.Read(bs); err != nil {
				log.Fatalf("Reading #%d failed: %v", i, err)
			}
		}(i)
	}

	defer func() {
		for i := 0; i < n; i++ {
			rl[i].Close()
			wl[i].Close()
		}
	}()

	for i := 0; i < n; i++ {
		<-reading
	}

	for i := 0; i < n; i++ {
		if _, err := wl[i].Write([]byte("Hello, World")); err != nil {
			log.Fatalf("Write #%d failed: %v", i, err)
		}
	}
}
