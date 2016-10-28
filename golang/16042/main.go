package main

import (
	"encoding/json"
	"os"
)

func main() {
	json.NewEncoder(os.Stdout).Encode(struct{ M json.Marshaler }{}) //bad: panics, should print `{"M":null}`
}
