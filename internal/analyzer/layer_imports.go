package analyzer

import (
	"fmt"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/nsamartsev/go-layer-linter/internal/config"
	"github.com/nsamartsev/go-layer-linter/internal/utils"
)

func Run(projectDir string) {
	cfg := config.LoadConfig(".golint.yaml")

	layerToPath := map[string]string{
		"handler":    "internal/handler",
		"service":    "internal/service",
		"repository": "internal/repository",
	}

	for _, rule := range cfg.LayeredImports {
		layer := rule.Layer
		allowed := rule.Allows

		layerPath := layerToPath[layer]
		allowedPath := layerToPath[allowed]

		files := utils.WalkGoFiles(filepath.Join(projectDir, layerPath))

		for _, file := range files {
			fset := token.NewFileSet()
			node, err := parser.ParseFile(fset, file, nil, parser.ImportsOnly)
			if err != nil {
				continue
			}

			for _, imp := range node.Imports {
				pkgName := strings.Trim(imp.Path.Value, "\"")

				if strings.Contains(pkgName, "internal") && !strings.HasPrefix(pkgName, allowedPath) {
					fmt.Printf("[ERROR] Forbidden import: %s -> %s in %s\n", layer, pkgName, file)
				}
			}
		}
	}
}
