package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err=%v\n", err)
			if false {
				var buf = make([]byte, 10)
				runtime.Stack(buf, false)
				fmt.Printf("err=%v stack=%s\n", err, buf)
			}
		}
	}()

	if false {
		var tl *string
		fmt.Printf("tl: %v\n", *tl)
	}
	tick := time.Tick(1e9)
	for {
		<-tick
	}
}
