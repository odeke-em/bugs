package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	now := time.Now()
	max := time.Unix(math.MaxInt64, 0)
	if now.Before(now) {
		fmt.Println("Cool, not the end of the world yet")
	} else {
		fmt.Println("uh oh")
	}
	ts := []time.Time{now, max}
	for i, t := range ts {
	    fmt.Printf("#%d t.sec() %v t.nsec() %v\n", i, uint64(t.Sec()), uint64(t.NSec()))
	}
}

	

	
