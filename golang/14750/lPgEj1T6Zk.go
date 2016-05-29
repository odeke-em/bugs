package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func main() {
	b := []byte(`{"typ":"JWS","alg":"HS256","alg":"foo"}`)
	var h Header
	if err := json.Unmarshal(b, &h); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", h)
}
