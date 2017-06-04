package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	m.Lock()
	m.Unlock()

	fmt.Printf("done here!\n")
}
