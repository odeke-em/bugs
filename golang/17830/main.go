package main

import "fmt"

type astruct struct {
	a, b int
}

func main() {
	var iface interface{} = map[string]astruct{}
	var iface2 interface{} = []astruct{}
	fmt.Println(iface, iface2)
}
