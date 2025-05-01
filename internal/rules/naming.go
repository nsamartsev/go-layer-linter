package rules

import (
	"fmt"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/nsamartsev/go-layer-linter/internal/utils"
)

var defaultNamingPatternsMap = map[string]string{
	"repository.go": "Repository",
	"service.go":    "Service",
	"usecase.go":    "UseCase",
}

func CheckNaming(filePath string) []string {
	var issues []string

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return []string{fmt.Sprintf("File parsing error %s: %v", filePath, err)}
	}

	structNames := utils.GetStructNames(node)
	relPath, _ := filepath.Rel(filepath.Dir(filePath), filePath)

	for _, name := range structNames {
		for suffix, pattern := range defaultNamingPatternsMap {
			if strings.HasSuffix(filePath, suffix) && !strings.HasSuffix(name, pattern) {
				issues = append(issues, fmt.Sprintf(
					"[ERROR] %s must follow the *%s pattern\n\t→ %s", name, pattern, relPath))
			}
		}
	}

	return issues
}
