package exec_race_test

import (
	"bytes"
	"testing"
	"text/template"
)

func TestExecuteRace(t *testing.T) {
	tmpl := template.Must(template.New("testing").Parse("{{.Name}}"))

	waitCount := 0
	n := 120
	waiterChan := make(chan bool)
	for i := 0; i < n; i++ {
		waitCount += 1
		go func() {
			defer func() { waiterChan <- true }()
			b := new(bytes.Buffer)
			st := struct{ Name string }{Name: "Golang"}
			_ = tmpl.Execute(b, st)
		}()
	}

	for i := 0; i < n; i++ {
		<-waiterChan
	}
}
