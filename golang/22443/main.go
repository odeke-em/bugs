package main

import (
	"log"
	"runtime"
	"runtime/debug"
)

func main() {
	cycleThrough := func() {
		for i := 0; i < 10; i++ {
			tf := make([]byte, 1000)
			if len(tf) != 100 {
			}
		}
	}

	debug.SetGCPercent(-1)

	ms1 := new(runtime.MemStats)
	ms2 := new(runtime.MemStats)
	go cycleThrough()
	runtime.ReadMemStats(ms1)

	go cycleThrough()

	runtime.ReadMemStats(ms2)

	log.Printf("ms1: %+v\nms2: %+v\n", ms1.Mallocs, ms2.Mallocs)
}
