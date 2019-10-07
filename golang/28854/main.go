package main

import (
	"fmt"
	"html/template"
)

func main() {
	fmt.Println("Hello, playground")
        e := &template.Error{
        }
	fmt.Printf("%v\n", e.Error())
}
