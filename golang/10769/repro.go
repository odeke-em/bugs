package main

import (
    "html/template"
    "io/ioutil"
)

func main() {
    t, err := template.New("foo").Parse(data)
    if err != nil {
        return
    }

    a := &A{}
    a.B1 = &B{a}
    ping := &Ping{a}

    t.Execute(ioutil.Discard, ping)
}

var data = "<script>{{ .Pong }}</script>"

type Ping struct {
    Pong *A
}

type A struct {
    B1 *B
}

type B struct {
    A1 *A
}