package main

import "time"

func main() {
	t := struct {
		about      string
		before     map[string]uint
		update     map[string]int
		updateTime time.Time
		expect     map[string]int
	}{
		about:   "this one",
		updates: map[string]int{"gopher": 10},
	}

	if t != nil {
	}
}
