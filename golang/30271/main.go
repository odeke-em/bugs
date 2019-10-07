package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("now\t\t%v\n", now)
                tnow := time.Now()
		fmt.Printf("time.Now()\t%v\n", tnow)
                fmt.Printf("Diff: %s\n\n", tnow.Sub(now))
	}
}
