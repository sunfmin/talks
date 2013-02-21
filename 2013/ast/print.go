package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	astf, _ := parser.ParseFile(fset, "/Users/sunfmin/gopkg/src/github.com/sunfmin/talks/2013/ast/source.go", nil, 0)
	ast.Print(fset, astf)

}

// endmain OMIT
