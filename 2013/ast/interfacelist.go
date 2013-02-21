package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	astf, _ := parser.ParseFile(fset, "/Users/sunfmin/gopkg/src/github.com/sunfmin/talks/2013/ast/source.go", nil, 0)

	w := &walker{}
	ast.Walk(w, astf)
	fmt.Println(w.interfacelist)
}

type walker struct {
	fset             *token.FileSet
	interfacelist    []string
	lastTypeSpecName string
}

// Visit implements the ast.Visitor interface.
func (w *walker) Visit(node ast.Node) ast.Visitor {

	switch n := node.(type) {
	case *ast.TypeSpec:
		w.lastTypeSpecName = n.Name.Name
	case *ast.InterfaceType:
		w.interfacelist = append(w.interfacelist, w.lastTypeSpecName)
	}
	return w
}

// endmain OMIT
