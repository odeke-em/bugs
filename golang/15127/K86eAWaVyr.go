package main

import (
	"encoding/json"
	"fmt"
)

type Id struct {
	Id1 int `json:"id"`
	Id2 int `json:"id"`
}

func main() {
	str := `{"id":123}`
	id := Id{}

	err := json.Unmarshal([]byte(str), &id)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}
