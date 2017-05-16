package main

import "plugin"
import "fmt"

func main() {
	aPlug, _ := plugin.Open("a.so")
	aSymPlug, _ := aPlug.Lookup("Rule")
	fmt.Printf("Plugin: %v loaded\n", aSymPlug)

	bPlug, _ := plugin.Open("b.so")
	bSymPlug, _ := bPlug.Lookup("Rule")
	fmt.Printf("Plugin: %v loaded\n", bSymPlug)
}
