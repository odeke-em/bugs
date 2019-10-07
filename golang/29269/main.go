package main

import (
	"html/template"
	"os"
)

func main() {
	t := template.Must(template.New("test").Parse(`
	{{- /* for some reason, 2 if blocks appear to be required to trigger this error */}}
	{{- /* not having 2 ifs, triggers another ambigous error */}}
	{{- /* doesn't matter what we check here */}}
    {{- if . }}
      <a href="mailto:{{.}}"></a>
    {{- end}}
	{{- /* nor here */}}
    {{- if .}}
    {{- /* the error happens here */}}
      <a href="{{.}}"></a>
    {{end}}`))

	// not that the style tag was not closed, this doesn't appear to be
	// happenening with any other tags
	t = template.Must(t.New("base").Parse(`<style>{{template "test" .}}`))

	// the error is reproduced with any data
	err := t.Execute(os.Stdout, "test")
	if err != nil {
		panic(err)
	}
}
