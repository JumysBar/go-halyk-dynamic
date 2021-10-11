package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "reflection/reflect_type.go", nil, 0)
	if err != nil {
		panic(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		switch s := n.(type) {
		case *ast.Ident:
			fmt.Printf("Identifier: %s\n", s)
		}
		return true
	})
}
