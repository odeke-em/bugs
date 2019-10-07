package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	pid := os.Getpid()
	ownProc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("Failed to find own process by pid %d error: %v", pid, err)
		os.Exit(1)
	}

	fmt.Printf("Pid: %d\n", pid)
	for {
		t1 := time.Now()
		time.Sleep(3 * time.Second)
		if true {
			ownProc.Signal(syscall.SIGTSTP)
		}
		t2 := time.Now()
		mt1, mt2 := monotonic(t1), monotonic(t2)
		td1, td2 := t1.UnixNano(), t2.UnixNano()
		diffTSec := float64(td2-td1) / 1e9
		diffMSec := float64(mt2-mt1) / 1e9
		fmt.Printf("Before: %d\nAfter:  %d\nDiff:   %f\nm1:     %d\nm2:     %d\ndiff:   %f%%\n"+
			"dT-dM:  %f\n\n", td1, td2, diffTSec, mt1, mt2, diffMSec, (100 * (diffMSec-diffTSec))/diffTSec)
	}
}

type captureMonotonic struct {
	_   uint64
	ext int64
}

func monotonic(t time.Time) int64 {
	return time.Mono(t)

	ct := (*captureMonotonic)(unsafe.Pointer(&t))
	return ct.ext
}
