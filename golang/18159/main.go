package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
)

const (
	templateString = `<script type="application/json">{{json .}}</script>`
	expectedString = `<script type="application/json">{"name":"Peter"}</script>`
)

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"json": func(obj interface{}) (template.JS, error) {
			data, err := json.Marshal(obj)
			if err != nil {
				return "", err
			}
			return template.JS(data), nil
		},
	}
}

type testPerson struct {
	Name string `json:"name"`
}

func main() {
	tmpl := template.New("").Funcs(templateFuncMap())
	if _, err := tmpl.Parse(templateString); err != nil {
		log.Fatalf("Could not parse template '%s': %s", templateString, err.Error())
	}
	buf := bytes.NewBuffer(nil)
	tmpl.Execute(buf, testPerson{Name: "Peter"})
	out := buf.String()
	if out != expectedString {
		log.Fatalf("Strings do not match: got '%s', want '%s'", out, expectedString)
	}
	log.Println("Works fine with this version of Go")
}
