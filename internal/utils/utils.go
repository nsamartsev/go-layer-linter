package utils

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

func GetPackageName(path string) (string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		return "", err
	}
	return file.Name.Name, nil
}

func GetStructNames(file *ast.File) []string {
	var names []string
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if _, ok := typeSpec.Type.(*ast.StructType); ok {
						names = append(names, typeSpec.Name.Name)
					}
				}
			}
		}
	}
	return names
}

func GetDirName(path string) string {
	return strings.TrimPrefix(filepath.Dir(path), "internal"+string(filepath.Separator))
}

func GetStructNamesFromFile(filePath string) []string {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil
	}

	var names []string
	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if _, ok := typeSpec.Type.(*ast.StructType); ok {
						names = append(names, typeSpec.Name.Name)
					}
				}
			}
		}
	}
	return names
}
