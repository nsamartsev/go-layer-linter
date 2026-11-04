// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package rules

import (
	"fmt"
	"go/ast"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/nsamartsev/go-layer-linter/internal/config"
)

func checkLayerImports(relativePath string, file *ast.File, cfg *config.Config, modulePath string) []string {
	currentLayer, currentPackage := layerForDirectory(path.Dir(relativePath), cfg.DDD.Layers)
	if currentLayer == "" {
		return nil
	}

	var issues []string
	for _, imported := range file.Imports {
		importPath, err := strconv.Unquote(imported.Path.Value)
		if err != nil || !importIsRelevantLayerImport(importPath, modulePath, cfg.DDD.Layers) {
			continue
		}

		allowed := layerImportMatches(importPath, currentPackage, modulePath)
		for _, allowedPackage := range cfg.DDD.Layers[currentLayer].ImportsAllowed {
			if layerImportMatches(importPath, allowedPackage, modulePath) {
				allowed = true
				break
			}
		}
		if !allowed {
			issues = append(issues, fmt.Sprintf("[ERROR] forbidden layer import: %s imports %s (%s)", currentLayer, importPath, relativePath))
		}
	}

	return issues
}

func importIsRelevantLayerImport(importPath, modulePath string, layers map[string]config.Layer) bool {
	if importIsProjectInternal(importPath, modulePath) {
		return true
	}
	for _, layer := range layers {
		if layerImportMatches(importPath, layer.Package, modulePath) {
			return true
		}
	}
	return false
}

func layerForDirectory(directory string, layers map[string]config.Layer) (string, string) {
	names := make([]string, 0, len(layers))
	for name := range layers {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		return len(layers[names[i]].Package) > len(layers[names[j]].Package)
	})

	for _, name := range names {
		packagePath := strings.Trim(layers[name].Package, "/")
		if isWithinPath(directory, packagePath) {
			return name, packagePath
		}
	}
	return "", ""
}
