package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func maybeStartProfile() func() {
	v := os.Getenv("CPUPROFILE")
	if v == "" {
		return func() {}
	}
	f, err := os.OpenFile(v, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("Failed to initiate profiling: %s. Ignoring", err)
		return func() {}
	}
	pprof.StartCPUProfile(f)
	return func() {
		pprof.StopCPUProfile()
		_ = f.Close()
	}
}

func routine1() {
	for {
	}
}

func routine2() {
	for {
	}
}

func main() {
	defer maybeStartProfile()()

	go routine1()
	go routine2()

	time.Sleep(30 * time.Second)
}