package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t, _ := template.New("test").Parse(`{{ (.).poo }}`)

	var v interface{}
	err := t.Execute(os.Stdout, v)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
}
