package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a reflect.Value
	switch a.(type) {
	case *reflect.InterfaceValue,
		*reflect.PtrValue:
		if v.IsNil() {
			fmt.Printf("null")
		}
	}
}
