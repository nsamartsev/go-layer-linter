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

func checkStructLocation(relativePath string, file *ast.File, cfg *config.Config) []string {
	var issues []string
	directory := path.Dir(relativePath)
	domainLayer, hasDomainLayer := cfg.DDD.Layers["domain"]
	domainPackage := strings.Trim(domainLayer.Package, "/")

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

			name := typeSpec.Name.Name
			if cfg.Rules.Locations.EntityInDomain && hasDomainLayer && strings.HasSuffix(name, "Entity") && !isWithinPath(directory, domainPackage) {
				issues = append(issues, fmt.Sprintf("[ERROR] entity %s must be declared under %s (%s)", name, domainPackage, relativePath))
			}
			_, isInterface := typeSpec.Type.(*ast.InterfaceType)
			if cfg.Rules.Locations.RepositoryInterfaceInDomainRepo && hasDomainLayer && isInterface && strings.HasSuffix(name, "Repository") && !isWithinPath(directory, path.Join(domainPackage, "repository")) {
				issues = append(issues, fmt.Sprintf("[ERROR] repository %s must be declared in %s (%s)", name, path.Join(domainPackage, "repository"), relativePath))
			}
			if cfg.Rules.Locations.UseCaseInUseCaseFile && strings.HasSuffix(name, "UseCase") && !isUseCaseFile(path.Base(relativePath)) {
				issues = append(issues, fmt.Sprintf("[ERROR] use case %s must be declared in a *_usecase.go file (%s)", name, relativePath))
			}
		}
	}

	return issues
}

func isUseCaseFile(fileName string) bool {
	return fileName == "usecase.go" || strings.HasSuffix(fileName, "_usecase.go")
}
