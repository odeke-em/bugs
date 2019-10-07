package main

import "sync"

func withMap(m map[string]int, i int) {
	if i&1 == 0 {
		m["key"] = 10
	} else {
		v := m["key"]
		if v == 0 {
			m["key"] = 11
		}
	}
}

func main() {
	var waitForAllGs sync.WaitGroup
	defer waitForAllGs.Wait()

	// A data race occurs when there are concurrent accesses
	// to shared data and at least one of the accesses is a write.
	shared := make(map[string]int)

	n := 10
	pauseTillAccess := make(chan bool)

	for i := 0; i < n; i++ {
		waitForAllGs.Add(1)
		go func(id int) {
			defer waitForAllGs.Done()
			pauseTillAccess <- true
			withMap(shared, id)
		}(i)
	}

	for i := 0; i < n; i++ {
		<-pauseTillAccess
	}
}
