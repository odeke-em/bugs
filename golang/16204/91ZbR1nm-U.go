package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
)

func main() {
	var buff bytes.Buffer

	var mp *map[int]int
	var ip *int
	ivs := []interface{}{nil, mp, ip}
	if true {
		for i, iv := range ivs {
			fmt.Printf("#%d: \n", i)
			gob.NewEncoder(&buff).Encode(iv)
		}
	}
	fmt.Println("reflect ", reflect.TypeOf(nil))
}
