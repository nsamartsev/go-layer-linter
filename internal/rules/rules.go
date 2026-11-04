// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package rules

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nsamartsev/go-layer-linter/internal/config"
)

// RunAnalysis analyzes Go files beneath projectDir using its go.mod and .dddlint.yaml files.
func RunAnalysis(projectDir string) ([]string, error) {
	root, err := filepath.Abs(projectDir)
	if err != nil {
		return nil, fmt.Errorf("resolve project directory %q: %w", projectDir, err)
	}
	info, err := os.Stat(root)
	if err != nil {
		return nil, fmt.Errorf("open project directory %s: %w", root, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("project path %s is not a directory", root)
	}

	cfg, err := config.LoadConfig(root)
	if err != nil {
		return nil, err
	}
	modulePath, err := readModulePath(root)
	if err != nil {
		return nil, fmt.Errorf("read Go module for project %s: %w", root, err)
	}

	var issues []string
	fset := token.NewFileSet()
	err = filepath.WalkDir(root, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if filePath != root {
				switch entry.Name() {
				case ".git", "vendor", "testdata", "examples", "fixtures", "samples":
					return filepath.SkipDir
				}
			}
			return nil
		}
		if filepath.Ext(filePath) != ".go" {
			return nil
		}

		relativePath, err := filepath.Rel(root, filePath)
		if err != nil {
			return fmt.Errorf("make path relative for %s: %w", filePath, err)
		}
		relativePath = filepath.ToSlash(relativePath)

		file, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
		if err != nil {
			return fmt.Errorf("parse %s: %w", relativePath, err)
		}

		issues = append(issues, checkLayerImports(relativePath, file, cfg, modulePath)...)
		issues = append(issues, checkNaming(relativePath, file, cfg)...)
		issues = append(issues, checkStructLocation(relativePath, file, cfg)...)
		issues = append(issues, checkForbiddenPackages(relativePath, file, cfg)...)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("analyze project %s: %w", root, err)
	}

	return issues, nil
}

func readModulePath(projectDir string) (string, error) {
	goModPath := filepath.Join(projectDir, "go.mod")
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", goModPath, err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(strings.TrimSpace(line))
		if len(fields) == 0 || fields[0] != "module" {
			continue
		}
		if len(fields) < 2 || strings.HasPrefix(fields[1], "//") {
			return "", fmt.Errorf("module directive in %s has no module path", goModPath)
		}

		modulePath := fields[1]
		if strings.HasPrefix(modulePath, "\"") || strings.HasPrefix(modulePath, "`") {
			modulePath, err = strconv.Unquote(modulePath)
			if err != nil {
				return "", fmt.Errorf("parse module path in %s: %w", goModPath, err)
			}
		}
		if modulePath == "" {
			return "", fmt.Errorf("module path in %s must not be empty", goModPath)
		}
		return modulePath, nil
	}

	return "", fmt.Errorf("%s has no module directive", goModPath)
}

func isWithinPath(candidate, parent string) bool {
	candidate = strings.Trim(filepath.ToSlash(filepath.Clean(candidate)), "/")
	parent = strings.Trim(filepath.ToSlash(filepath.Clean(parent)), "/")
	return candidate == parent || strings.HasPrefix(candidate, parent+"/")
}

func packagePathMatches(importPath, configuredPath string) bool {
	importPath = strings.Trim(filepath.ToSlash(importPath), "/")
	configuredPath = strings.Trim(filepath.ToSlash(filepath.Clean(strings.TrimSpace(configuredPath))), "/")
	if configuredPath == "" || configuredPath == "." {
		return false
	}
	return importPath == configuredPath || strings.HasPrefix(importPath, configuredPath+"/")
}

func layerImportMatches(importPath, configuredPath, modulePath string) bool {
	if packagePathMatches(importPath, configuredPath) {
		return true
	}

	modulePath = strings.Trim(filepath.ToSlash(filepath.Clean(strings.TrimSpace(modulePath))), "/")
	if modulePath == "" || modulePath == "." {
		return false
	}
	configuredPath = strings.Trim(filepath.ToSlash(filepath.Clean(strings.TrimSpace(configuredPath))), "/")
	if configuredPath == "" || configuredPath == "." {
		return false
	}
	fullPath := modulePath + "/" + configuredPath
	return packagePathMatches(importPath, fullPath)
}

func importIsProjectInternal(importPath, modulePath string) bool {
	modulePath = strings.Trim(filepath.ToSlash(filepath.Clean(strings.TrimSpace(modulePath))), "/")
	if modulePath == "" || modulePath == "." {
		return false
	}
	modulePrefix := modulePath + "/"
	if !strings.HasPrefix(importPath, modulePrefix) {
		return false
	}
	return importIsInternal(strings.TrimPrefix(importPath, modulePrefix))
}

func importIsInternal(importPath string) bool {
	return importPath == "internal" || strings.HasPrefix(importPath, "internal/") || strings.Contains(importPath, "/internal/")
}
