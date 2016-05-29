package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"runtime"
)

func caller() string {
	_, p, _, _ := runtime.Caller(1)
	return p
}

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	// Parse the file containing this very example
	// but stop after processing the imports.
	fmt.Println(caller())
	f, err := parser.ParseFile(fset, caller(), nil, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the imports from the file's AST.
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}

}