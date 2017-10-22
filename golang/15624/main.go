package main

import (
	"encoding/json"
	"log"
)

var jsdata = []byte(`{ "y": ["1", "2", "3", "3333333333333333211", 2999], "B": [1, 3]}`)

type foo struct {
	Y []int64
}

type bar struct {
	Y []int64 `json:",string"`
	B []int64
}

func main() {
	tests := []interface{}{
		new(bar),
		new(foo),
	}
	for i, recv := range tests {
		if err := json.Unmarshal(jsdata, recv); err != nil {
			log.Printf("err:: #%d: %v", i, err)
		} else {
			log.Printf("pass: #%d: %#v\n", i, recv)
		}
	}
}
