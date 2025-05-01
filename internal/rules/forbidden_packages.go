package rules

import (
	"fmt"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/nsamartsev/go-layer-linter/internal/config"
	"github.com/nsamartsev/go-layer-linter/internal/utils"
)

func CheckForbiddenPackages(filePath string, cfg *config.Config) []string {
	var issues []string
	dir := utils.GetDirName(filePath)

	if !strings.Contains(dir, "domain") {
		return nil // только для domain
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return []string{fmt.Sprintf("Ошибка парсинга %s: %v", filePath, err)}
	}

	for _, imp := range node.Imports {
		path := strings.Trim(imp.Path.Value, "\"")

		// Пропускаем импорты из internal
		if strings.HasPrefix(path, "internal/") {
			continue
		}

		// Проверяем, есть ли такой пакет в списке запрещённых
		for _, forbidden := range cfg.ForbiddenPackagesInDomain {
			if path == forbidden {
				relPath, _ := filepath.Rel(dir, filePath)
				issues = append(issues, fmt.Sprintf(
					"[ERROR] Forbidden import in domain: %s\n\t→ %s",
					path, relPath))
			}
		}
	}

	return issues
}
