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
	for _, i := range w.interfacelist {
		fmt.Println("interface: ", i.Name)
		for _, m := range i.Methods {
			fmt.Println("\tmethod: ", m.Name)
		}
	}
}

type Interface struct {
	Name    string
	Methods []*Method
}

type Method struct {
	Name string
}

type walker struct {
	fset             *token.FileSet
	interfacelist    []*Interface
	lastTypeSpecName string
}

// Visit implements the ast.Visitor interface.
func (w *walker) Visit(node ast.Node) ast.Visitor {

	switch n := node.(type) {
	case *ast.TypeSpec:
		w.lastTypeSpecName = n.Name.Name
	case *ast.InterfaceType:
		i := &Interface{Name: w.lastTypeSpecName}
		w.interfacelist = append(w.interfacelist, i)
		ms := n.Methods.List
		for _, m := range ms {
			i.Methods = append(i.Methods, &Method{Name: m.Names[0].Name})
		}
	}
	return w
}

// endmain OMIT
