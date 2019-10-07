package main

import (
	"flag"
	"log"
	"os"
	"runtime/debug"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
        var n int
        var th int
	flag.IntVar(&n, "n", 100, "number of goroutines")
	flag.IntVar(&th, "t", 10, "number of threads")
	flag.Parse()
	f, err := os.Create("go1.13.txt")
	if err != nil {
		log.Fatalf("Failed to create trace file: %v", err)
	}
	defer f.Close()

        println("n", n, "threads", th)
	var rl, wl []*os.File
	cr := make(chan bool, n)

	for i := 0; i < n; i++ {
		fpr, fpw, err := os.Pipe()
		if err != nil {
			for j := 0; j < i; i++ {
				fpr.Close()
				fpw.Close()
			}
			log.Fatalf("Failed to create os.Pipe: %v", err)
		}
		wl = append(wl, fpw)
		rl = append(rl, fpr)

		defer fpr.Close()
		defer fpw.Close()
	}

	trace.Start(f)
	defer trace.Stop()

	defer debug.SetMaxThreads(debug.SetMaxThreads(th))

	var wg sync.WaitGroup
	defer wg.Wait()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cr <- true
			var b [4]byte
			fpr := rl[id]
			if _, err := fpr.Read(b[:]); err != nil {
				log.Fatalf("Read failed: %v", err)
			}
		}(i)
	}

	for i := 0; i < n; i++ {
		<-cr
	}

	<-time.After(3 * time.Second)

	for i := 0; i < n; i++ {
		fpw := wl[i]
		if _, err := fpw.Write([]byte("Hello")); err != nil {
			log.Fatalf("Failed to write: %v", err)
		}
	}
}
