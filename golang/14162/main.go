package main

import (
	"fmt"
	"runtime"
)

var memStats = new(runtime.MemStats)

func AllocatableMemory() uint64 {
	return memStats.Alloc
}

func main() {
	prevAlloc := uint64(0)
	for i := 0; i < 10000; i++ {
		runtime.ReadMemStats(memStats)
		vx := new(runtime.MemStats)
		if memStats.HeapSys != prevAlloc {
			fmt.Printf("\033[31mSpike #%d: old: %d new: %d\033[00m\n", i, prevAlloc, memStats.HeapSys)
			prevAlloc = memStats.HeapSys
		fmt.Printf("#%d: heapSys: %d alloc: %d heapAlloc: %d\n", i, memStats.HeapSys, memStats.Alloc, memStats.HeapAlloc)
		}
		vx = nil
		if i%2 == 3 {
			fmt.Printf("vx=%p\n", vx)
		}
	}
}
