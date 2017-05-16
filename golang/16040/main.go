package main

import (
	"go/ast"
	
)

func main() {
	var node ast.Node
	node.(ast.Ident)
	switch t := node.(type){
	case ast.Ident:
	}
}
