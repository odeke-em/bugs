// +build ignore

package main

import (
	"bytes"
	"fmt"
	"html/template"
)

func main() {
	all := []string{"title.xhtml", "config", "xhtml"}
	count := make(map[string]int)
	for i := 0; i < 100; i++ {
		var res *template.Template
		for _, name := range all {
			var tmpl *template.Template
			if res == nil {
				tmpl = template.New(name)
				res = tmpl
			} else {
				tmpl = res.New(name)
			}
			_, err := tmpl.Parse(inlined[name])
			fmt.Printf("#%d: err: %v\n", i, err)
		}

		var buf bytes.Buffer
		res.Execute(&buf, map[string]interface{}{
			"Book": map[string]interface{}{},
		})
		count[buf.String()] += 1
	}
	for output, n := range count {
		fmt.Println(n, output, "---")
	}
}

var inlined = map[string]string{
	"title.xhtml": `{{template "xhtml" .}}`,
	"config": `{{define "stylesheets" -}}
<link rel="stylesheet" type="text/css" href="{{.Book.CSSPath}}"/>
{{end -}}`,
	"xhtml": `<!DOCTYPE html>
<html>
<head>
{{block "stylesheets" .}}{{end -}}
</head>
`,
}
