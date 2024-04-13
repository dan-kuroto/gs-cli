package utils

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strconv"
)

func AddImport(fpath string, importPath string) string {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fpath, nil, parser.ParseComments)
	if err != nil {
		ThrowE(err)
	}

	added := false
	for _, decl := range file.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			if !added && decl.Tok == token.IMPORT {
				added = true
				decl.Specs = append(decl.Specs, &ast.ImportSpec{
					Name: &ast.Ident{
						Name: "_",
					},
					Path: &ast.BasicLit{
						Value: strconv.Quote(importPath),
					},
				})
			}
		}
	}

	ast.SortImports(fset, file)

	outStr, err := GenerateCode(fset, file)
	if err != nil {
		ThrowE(err)
	}
	return outStr
}

func GenerateCode(fset *token.FileSet, file *ast.File) (string, error) {
	var buffer bytes.Buffer
	if err := printer.Fprint(&buffer, fset, file); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
