package main

import (
	"bytes"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	prc, pwc, err := os.Pipe()
	if err != nil {
		log.Fatalf("Failed to create os.Pipe: %v", err)
	}

	f, err := os.Create("trace.txt")
	if err != nil {
		log.Fatalf("Failed to create trace.txt file: %v", err)
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	readyToRead := make(chan bool)
	readyToWrite := make(chan bool)
	go func() {
                <-readyToRead
		<-readyToWrite
		pwc.Write(bytes.Repeat([]byte("A"), 200))
	}()

	resCh := make(chan int)
	go func() {
		var n int
		defer func() {
			// Intentionally putting readyToWrite here so that
			// we can block until the process is killed.
			close(readyToWrite)
			resCh <- n
		}()

		bs := make([]byte, 100)
		var err error
		close(readyToRead)
		n, err = prc.Read(bs)
		if err != nil {
			log.Fatalf("Failed to perform read: %v", err)
		}
		log.Printf("Read: %d\n", n)
	}()

	<-readyToRead

	select {
	case <-time.After(10 * time.Second):
            panic("Exiting")
		// Good bye.
	case <-resCh:
	}
}
