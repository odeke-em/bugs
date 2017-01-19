package main

import "plugin"

// you can import nearly any package (tested with fmt, log, json, etc.)
import "encoding/json"
var dummy json.RawMessage

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}

	f.(func())()
}