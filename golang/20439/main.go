package main

import (
	"bytes"
	"fmt"
	"text/template"
)

var tmpl = template.Must(template.New("").Parse("Element 0 is {{index . 1}}"))

func applyTemplate(v interface{}) {
	buf := bytes.NewBuffer(nil)
	err := tmpl.Execute(buf, v)
	if err != nil {
		fmt.Printf("[err] %v\n", err)
	} else {
		fmt.Printf("[ ok] %v\n", buf)
	}
}

func main() {
	applyTemplate([]string{"zero", "one", "two"})
	applyTemplate(map[int]string{0: "zero", 1: "one", 2: "two"})
	applyTemplate(map[int32]string{0: "zero", 1: "one", 2: "two"})
}
