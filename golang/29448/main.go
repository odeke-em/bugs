package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"testing"
)

func TestProfileSignal(t *testing.T) {
    buf := new(bytes.Buffer)

    pprof.StartCPUProfile(buf)
    go func() {
        select {}
    }()

    buf := make([]byte, 9999)
}
