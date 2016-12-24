package main

import (
	"fmt"
	"html/template"
	"os"
)

type CCV struct {
	Sections []Section
}

type Section struct {
	label string
}

//List returns sections that belong to block with "label" or empty section if it fails
func (ccv *CCV) List(label string) []Section {
	return ccv.Sections
}

const templ3 = `
Research
Grants
{{range Block "Research Funding History"}}
	{{.Emit "Funding Start Date, Funding End Date, Funding Title" }}
	{{range .SubSection "Funding Sources" }}
		{{.Emit "Program Name, Funding Organization" }}
	{{end}}	
{{end}}
`

func main() {
	ccv := &CCV{}
	tm, err := template.New("test").Funcs(template.FuncMap{"List": ccv.List}).Parse(templ3)
	if err != nil {
		fmt.Printf("Parsing failed: %s", err)
	}
	err = tm.Execute(os.Stdout, ccv)
	if err != nil {
		fmt.Printf("execution failed: %s", err)
	}

	fmt.Println("Hello, playground")
}
