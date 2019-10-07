package Äfoo

import (
	"fmt"
	"runtime"
)

var ÄbarV int = 101

func Äbar(x int) int {
	defer func() { ÄbarV += 3 }()
	return Äblix(x)
}

func Äblix(x int) int {
	defer func() { ÄbarV += 9 }()
	PrintTrace()
	return ÄbarV + x
}

func PrintTrace() {
	// Ask runtime.Callers for up to 10 pcs, including runtime.Callers itself.
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	if n == 0 {
		// No pcs available. Stop now.
		// This can happen if the first argument to runtime.Callers is large.
		return
	}

	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)

	// Loop to get frames.
	// A fixed number of pcs can expand to an indefinite number of Frames.
	for {
		frame, more := frames.Next()
		fmt.Printf("- more:%v | %s\n", more, frame.Function)
		if !more {
			break
		}
	}
}
