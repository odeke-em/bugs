package main

import (
	"plugin"
)

func main() {
	p, err := plugin.Open("post.so")
	if err != nil {
		panic(err)
	}

	add, err := p.Lookup("Init")
	if err != nil {
		panic(err)
	}
	add.(func())()
}