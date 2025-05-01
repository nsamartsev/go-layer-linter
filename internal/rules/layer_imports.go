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

func CheckLayerImports(filePath string, cfg *config.Config) []string {
	var issues []string

	currentDir := utils.GetDirName(filePath)

	var currentLayerName string

	fmt.Println(cfg.Layers)

	for name, layer := range cfg.Layers {
		fmt.Println("name=", name, "layer=", layer)
		if strings.HasPrefix(currentDir, filepath.Base(layer.Package)) {
			currentLayerName = name
			break
		}
	}

	if currentLayerName == "" {
		return nil // вне DDD-слоев
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return []string{fmt.Sprintf("Ошибка парсинга %s: %v", filePath, err)}
	}

	for _, imp := range node.Imports {
		path := strings.Trim(imp.Path.Value, "\"")

		if !strings.HasPrefix(path, "internal/") {
			continue
		}

		impDir := filepath.Base(filepath.Dir(path))

		allowed := false
		for _, allowedImport := range cfg.Layers[currentLayerName].ImportsAllowed {
			if impDir == filepath.Base(allowedImport) {
				allowed = true
				break
			}
		}

		if !allowed {
			issues = append(issues, fmt.Sprintf("[ERROR] Запрещён импорт: %s -> %s в файле %s", currentLayerName, impDir, filePath))
		}
	}

	return issues
}
