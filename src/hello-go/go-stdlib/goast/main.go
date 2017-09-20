package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	// ast.Print(fset, f)
	_ = f
	inspect()
}

func inspect() {
	expr, _ := parser.ParseExpr(`v+1`)
	ast.Inspect(expr, func(n ast.Node) bool {
		if n != nil {
			fmt.Printf("%T\n", n)
		}
		return true
	})
}
