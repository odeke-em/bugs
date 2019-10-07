package main

import (
	"bytes"
	"log"
	"strings"
	"html/template"
)

type F struct {
	name string
}

func (f *F) Name() string {
	return f.name
}

func (f *F) Other() *F {
	return nil
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

	result := strings.TrimSpace(buf.String())

	if !strings.Contains(result, "foo") {
		log.Fatal(result)
	}

}
