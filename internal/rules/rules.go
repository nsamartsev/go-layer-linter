package rules

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nsamartsev/go-layer-linter/internal/config"
)

func RunAnalysis(projectDir string) []string {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Loading of config file error: %v\n", err)
		os.Exit(1)
	}

	var issues []string

	filepath.Walk(projectDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			issues = append(issues, CheckLayerImports(path, cfg)...)
			issues = append(issues, CheckNaming(path)...)
			issues = append(issues, CheckStructLocation(path)...)
			issues = append(issues, CheckForbiddenPackages(path, cfg)...)
		}

		return nil
	})

	return issues
}
