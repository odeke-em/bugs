package main

import "text/template"

func main() {
	tmpl, _ := template.New("t").Funcs(template.FuncMap{"f": func(s string) string { return s }}).Parse("{{ . | f }}")
	if err := tmpl.Execute(nil, nil); err != nil {
		panic(err)
	}
}
