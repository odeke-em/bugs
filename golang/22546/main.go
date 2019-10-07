package main

import (
	"encoding/json"
)

func main() {
	b := []byte(`{"ID": "1"}`)
	var m Manager
	json.Unmarshal(b, &m)
	println(m.ID)
}

type Manager struct {
	*employee
}

type employee struct {
	ID string
}