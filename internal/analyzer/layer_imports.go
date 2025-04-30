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

	for layer, allowed := range cfg.LayeredImports {
		allowedPath := layerToPath[allowed]
		layerPath := layerToPath[layer]

		files := utils.WalkGoFiles(filepath.Join(projectDir, layerPath))

		for _, file := range files {
			fset := token.NewFileSet()
			node, err := parser.ParseFile(fset, file, nil, parser.ImportsOnly)
			if err != nil {
				continue
			}

			for _, imp := range node.Imports {
				pkgName := strings.Trim(imp.Path.Value, "\"")

				if !strings.Contains(pkgName, allowedPath) && strings.Contains(pkgName, "internal") {
					fmt.Printf("[ERROR] Forbidden import: %s -> %s in %s\n", layer, pkgName, file)
				}
			}
		}
	}
}
