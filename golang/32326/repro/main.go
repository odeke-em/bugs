package main

import (
	"log"
	"os"
	"runtime/debug"
	"sync"
	"syscall"
	"time"
)

func main() {
	n := 100

	var rl, wl []*os.File
	for i := 0; i < n; i++ {
		prc, pwc, err := os.Pipe()
		if err != nil {
			// Cleanup.
			for j := 0; j < i; j++ {
				wl[i].Close()
				rl[i].Close()
			}
		}
		rl = append(rl, prc)
		wl = append(rl, pwc)
	}

	defer debug.SetMaxThreads(debug.SetMaxThreads(10))

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			bs := make([]byte, 4)
			wg.Done()
			prc := rl[i]
			if _, err := prc.Read(bs); err != nil {
				log.Fatalf("Failed to read %d: %v\n", i, err)
			}
		}(i)
	}

	wg.Wait()
	println("Waiting now")
	for {
		<-time.After(5 * time.Second)
		if true {
			return
		}
		proc, _ := os.FindProcess(os.Getpid())
		proc.Signal(syscall.SIGQUIT)
	}

	for i := 0; i < n; i++ {
		if _, err := wl[i].Write([]byte("Hello")); err != nil {
			log.Fatalf("Write #%d failed: %v", err)
		}
	}
}
