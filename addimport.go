package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"strconv"
	//"github.com/josephspurrier/apigen/spec/user"
)

func AddImportToFile(file string) {
	// Create the AST by parsing src
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, 0)
	if err != nil {
		log.Println(err)
		return
	}

	// Import declaration.
	// Source: https://golang.org/src/go/doc/example.go#L262
	/*importDecl := &ast.GenDecl{
		Tok:    token.IMPORT,
		Lparen: 1, // Need non-zero Lparen and Rparen so that printer
		Rparen: 1, // treats this as a factored import.
	}*/

	// Copy over the old imports
	/*for _, s := range f.Imports {
		iSpec := &ast.ImportSpec{Path: &ast.BasicLit{Value: s.Path.Value}}
		importDecl.Specs = append(importDecl.Specs, iSpec)
	}*/

	// Add the new import
	//iSpec := &ast.ImportSpec{Path: &ast.BasicLit{Value: strconv.Quote("ast")}}
	//importDecl.Specs = append(importDecl.Specs, iSpec)

	// Add the imports
	for i := 0; i < len(f.Decls); i++ {
		d := f.Decls[i]

		switch d.(type) {
		case *ast.FuncDecl:
			// No action
		case *ast.GenDecl:
			dd := d.(*ast.GenDecl)

			// IMPORT Declarations
			if dd.Tok == token.IMPORT {
				// Add the new import
				iSpec := &ast.ImportSpec{Path: &ast.BasicLit{Value: strconv.Quote("ast")}}
				dd.Specs = append(dd.Specs, iSpec)
			}
		}
	}

	// Sort the imports
	ast.SortImports(fset, f)

	// Generate the code
	out, err := GenerateFile(fset, f)
	if err != nil {
		log.Println(err)
		return
	}

	// Output the screen
	fmt.Println(string(out))
}

func GenerateFile(fset *token.FileSet, file *ast.File) ([]byte, error) {
	var output []byte
	buffer := bytes.NewBuffer(output)
	if err := printer.Fprint(buffer, fset, file); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func main() {
	// 1. 似乎修改自己的时候无效 2. GenerateFile函数不是创建文件而是得到目标文件的string
	AddImportToFile("main.go")
}
