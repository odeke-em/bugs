package main

import "plugin"
import "os"

func run(path string) {
	myplugin, _ := plugin.Open(path)
	symbol, _ := myplugin.Lookup("V")
	instance := *symbol.(*int)
	println(instance)
}

func main() {
	for _, arg := range os.Args[1:] {
		run(arg)
	}
}