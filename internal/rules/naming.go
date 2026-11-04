// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package rules

import (
	"fmt"
	"go/ast"
	"go/token"
	"path"
	"strings"

	"github.com/nsamartsev/go-layer-linter/internal/config"
)

func checkNaming(relativePath string, file *ast.File, cfg *config.Config) []string {
	fileName := path.Base(relativePath)
	patterns := []struct {
		fileSuffix string
		typeSuffix string
		enabled    bool
	}{
		{fileSuffix: "repository.go", typeSuffix: "Repository", enabled: cfg.Rules.Naming.Repository},
		{fileSuffix: "service.go", typeSuffix: "Service", enabled: cfg.Rules.Naming.Service},
		{fileSuffix: "usecase.go", typeSuffix: "UseCase", enabled: cfg.Rules.Naming.UseCase},
	}
	var issues []string
	for _, declaration := range file.Decls {
		genDecl, ok := declaration.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			for _, pattern := range patterns {
				if pattern.enabled && strings.HasSuffix(fileName, pattern.fileSuffix) && !strings.HasSuffix(typeSpec.Name.Name, pattern.typeSuffix) {
					issues = append(issues, fmt.Sprintf("[ERROR] %s must follow the *%s naming pattern (%s)", typeSpec.Name.Name, pattern.typeSuffix, relativePath))
				}
			}
		}
	}
	return issues
}
