package main

import "fmt"

func main() {
	people := map[string]*person{
		"a": &person{},
		"b": &person{},
		"c": &person{},
	}

	tip(people, "a", "c")
	tip(people, "b", "c", "b")
	tip(people, "a", "c", "c", "c", "c")

	for id, p := range people {
		fmt.Printf("Id=%q Person=%+v\n", id, p)
	}
}

type person struct {
	Age        float64
	Tips       float64
	Department int
}

func tip(people map[string]*person, ids ...string) {
	for _, id := range ids {
		if _, ok := people[id]; !ok {
			continue
		}
		if people[id].Tips >= 100 {
			continue
		}
		people[id].Tips += 1.2
		people[id].Age += 0.01
	}
}
