package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Field *string `xml:",comment"`
}

func drill(val reflect.Value) (reflect.Value, bool) {
	ok := true
	for val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		if val.IsNil() {
			ok = false
		}
		val = val.Elem()
	}
	return val, ok
}

func main() {
	av := &A{}
	rv := reflect.ValueOf(av)
	val, ok := drill(rv)
	fmt.Printf("val: %#v ok: %v\n", val, ok)
	switch val.Kind() {
	case reflect.Struct:
		fmt.Printf("isStruct")
	default:
		fmt.Printf("isOther")
	}
}