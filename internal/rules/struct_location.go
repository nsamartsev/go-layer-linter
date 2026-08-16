package rules

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/nsamartsev/go-layer-linter/internal/utils"
)

func CheckStructLocation(filePath string) []string {
	var issues []string

	dir := utils.GetDirName(filePath)
	fileName := filepath.Base(filePath)

	structNames := utils.GetStructNamesFromFile(filePath)

	for _, name := range structNames {

		relPath, _ := filepath.Rel(dir, filePath)

		if strings.Contains(name, "Entity") && !strings.HasPrefix(dir, "domain") {
			issues = append(issues, fmt.Sprintf(
				"[ERROR] Entity '%s' must be in  internal/domain directory\n\t→ %s",
				name,
				relPath,
			))
		}
		if strings.Contains(name, "Repository") && !strings.HasPrefix(dir, "repository") {
			issues = append(issues, fmt.Sprintf(
				"[ERROR] Repository '%s' must be in internal/domain/repository directory\n\t→ %s",
				name,
				relPath,
			))
		}
		if strings.Contains(name, "UseCase") && !strings.HasPrefix(fileName, "usecase") {
			issues = append(issues, fmt.Sprintf(
				"[ERROR] UseCase '%s' must be in file usecase_*.go\n\t→ %s",
				name,
				relPath,
			))
		}
	}

	return issues
}
