package main

import (
	"bytes"
	"log"
	"text/template"
)

type Namer interface {
	Name() string
}

type F struct {
	name string
}

func (f *F) Name() string {
	return f.name
}

func (f *F) Other() Namer {
	var n Namer
	return n
}

func main() {
	var buf bytes.Buffer
	tmpl, err := template.New("").Parse(`{{ .Other.Name }}`)
	if err != nil {
		log.Fatal(err)
	}

	data := &F{name: "foo"}

	if err := tmpl.Execute(&buf, data); err != nil {
		log.Fatal(err)
	}

}
