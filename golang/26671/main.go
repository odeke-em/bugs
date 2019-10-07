package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
)

var src = `
//line main.go:1
package p
`

func main() {
	tmpf, err := ioutil.TempFile("", "parser_test")
	if err != nil {
		panic(err)
	}
	defer tmpf.Close()

	err = ioutil.WriteFile(tmpf.Name(), []byte(src), os.ModePerm)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, tmpf.Name(), nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	properFileName := filepath.Join(filepath.Dir(tmpf.Name()), "main.go")
	fmt.Printf("File name is\n\t%q\nbut must be\n\t%q\n", fset.Position(f.Pos()).Filename, properFileName)
}
