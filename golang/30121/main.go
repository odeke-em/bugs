package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	var src = fmt.Sprintf(`
package main

func main() {
    var src = "%s"
    println(src)
}`, strings.Repeat("a", 1<<20))

// println(src)
	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
	if err != nil {
		log.Fatalf("parser.ParseFile failed: %v", err)
	}
	_ = f
	if err := printer.Fprint(ioutil.Discard, fset, new(ast.BlockStmt)); err != nil {
		log.Fatalf("printer.Fprint error: %v", err)
	}
}
