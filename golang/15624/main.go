package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var jsdata = []byte(`{ "A": ["1", "2", "3", "3333333333333333211", 2999] }`)

type foo struct {
	A []int64
}

func main() {
	var v foo
	if err := json.Unmarshal(jsdata, &v); err != nil {
		log.Fatal(err)
	}
	fmt.Println(v.A)
}
