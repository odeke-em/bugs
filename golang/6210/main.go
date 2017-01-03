package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	var before, after uint64
	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	before = m.Alloc
	if time.Now().Before(time.Unix(math.MaxInt64, 0)) {
	} else {
	}
	runtime.ReadMemStats(m)
	after = m.Alloc
	fmt.Printf("diff: %d\n", after-before)
}
