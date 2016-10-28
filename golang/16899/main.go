package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println(LoadConfig())
}

func LoadConfig() error {
	type config struct {
		Foo string
	}

	f, err := os.Open("foo.json")
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(f)
	cf := new(config)
	if err := decoder.Decode(cf); err != nil {
		return err
	}

	return nil
}
