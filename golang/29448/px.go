package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

const N = 10000

func main() {
	f, err := os.Create("/tmp/prof")
	if err != nil {
		log.Fatal(err)
	}

	pprof.StartCPUProfile(f)

	go func(){
		select {}
	}()

	buf := make([]byte, 99999)
	for i := 0; i < N; i++ {
		runtime.Stack(buf, true)
	}

	pprof.StopCPUProfile()
	f.Close()
}