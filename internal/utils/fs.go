package utils

import (
	"os"
	"path/filepath"
)

func WalkGoFiles(root string) []string {
	var goFiles []string

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			goFiles = append(goFiles, path)
		}
		return nil
	})

	return goFiles
}
