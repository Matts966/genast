package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("Please pass me filename 💀")
	}
	fileName := os.Args[1]
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, nil, 0)
	if err != nil {
		fmt.Printf("Error on parser.ParseFile: %v\n", err)
		return
	}
	ast.Print(fset, f)
}
