//+build !windows

package main

import (
	"log"
	"os"
	"syscall"
	"time"
)

func main() {
	pid := os.Getpid()
	log.Printf("PID: %d", pid)
	proc, err := os.FindProcess(pid)
	if err != nil {
		log.Fatalf("Failed to find own process with PID: %d err: %v", pid, err)
	}

	doneCh := make(chan bool)
	successCh := make(chan bool, 1)
	go func() {
		defer close(doneCh)

		for i := 0; i < 100; i++ {
			select {
			case <-successCh:
				return
			default:
				proc.Signal(syscall.SIGIO)
			}
		}
	}()
	<-time.After(2 * time.Second)
	successCh <- true
	<-doneCh
}
