package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type A struct {
	Field *string `xml:",comment"`
}

func main() {
	s := "<A><!--FieldValue--></A>"
	value := "FieldValue"
	msg1 := &A{}
	err := xml.Unmarshal([]byte(s), msg1)
	if err != nil {
		panic(err)
	}
	msg2 := &A{
		Field: &value,
	}
	data, err := xml.Marshal(msg2)
	if err != nil {
		panic(err) // go1.8rc3 panics here with:
		// panic: xml: bad type for comment field of main.A
	}
	if !bytes.Equal(data, []byte(s)) {
		panic(fmt.Errorf("expected %s got %s", s, string(data)))
	}
}
