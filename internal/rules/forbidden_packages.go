// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package rules

import (
	"fmt"
	"go/ast"
	"path"
	"strconv"

	"github.com/nsamartsev/go-layer-linter/internal/config"
)

func checkForbiddenPackages(relativePath string, file *ast.File, cfg *config.Config) []string {
	domainLayer, hasDomainLayer := cfg.DDD.Layers["domain"]
	if !hasDomainLayer {
		return nil
	}
	if !isWithinPath(path.Dir(relativePath), domainLayer.Package) {
		return nil
	}

	var issues []string
	for _, imported := range file.Imports {
		importPath, err := strconv.Unquote(imported.Path.Value)
		if err != nil {
			continue
		}
		for _, forbidden := range cfg.ForbiddenPackagesInDomain {
			if packagePathMatches(importPath, forbidden) {
				issues = append(issues, fmt.Sprintf("[ERROR] forbidden import in domain: %s (%s)", importPath, relativePath))
				break
			}
		}
	}

	return issues
}
