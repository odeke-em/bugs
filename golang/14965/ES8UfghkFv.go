package main

import (
	"fmt"
)

func main() {
	defer func() {
		i := recover()
		if i == nil {
			return
		}
		
		fmt.Printf("%v\n", i)
		if _, ok := i.(error); !ok {
			if _, ok := i.(string); !ok {
				fmt.Println("assignment to nil map panics with some other type")	
			} else {
				fmt.Println("assignment to nil map panics with type string")	
			}
		} else {
			fmt.Println("assignment to nil map panics with type error")
		}
	}()

	var theMap map[uint64]bool
	theMap[1234] = true
}
