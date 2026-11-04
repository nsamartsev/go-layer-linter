// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package rules

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const testConfig = `forbidden_packages_in_domain:
  - database/sql
  - net/http
ddd:
  layers:
    domain:
      package: internal/domain
      imports_allowed: []
    application:
      package: internal/application
      imports_allowed:
        - internal/domain
    infrastructure:
      package: internal/infrastructure
      imports_allowed:
        - internal/domain
        - internal/application
    interfaces:
      package: internal/interfaces
      imports_allowed:
        - internal/application
`

func TestRunAnalysis(t *testing.T) {
	tests := []struct {
		name        string
		files       map[string]string
		wantIssue   string
		wantNoIssue bool
	}{
		{
			name: "allows configured layers and subpackages",
			files: map[string]string{
				"internal/domain/user_entity.go": `package domain
type UserEntity struct{}
`,
				"internal/domain/repository/user_repository.go": `package repository
import "github.com/example/project/internal/domain"
type UserRepository interface { Find(domain.UserEntity) error }
`,
				"internal/application/user_usecase.go": `package application
import "github.com/example/project/internal/domain/repository"
type UserUseCase struct { repo repository.UserRepository }
`,
			},
			wantNoIssue: true,
		},
		{
			name: "reports forbidden layer import",
			files: map[string]string{
				"internal/application/user_usecase.go": `package application
import "github.com/example/project/internal/infrastructure"
type UserUseCase struct { repo infrastructure.Repository }
`,
			},
			wantIssue: "forbidden layer import",
		},
		{
			name: "reports forbidden domain package",
			files: map[string]string{
				"internal/domain/user_entity.go": `package domain
import "net/http"
type UserEntity struct { client *http.Client }
`,
			},
			wantIssue: "forbidden import in domain: net/http",
		},
		{
			name: "reports naming and location errors",
			files: map[string]string{
				"internal/application/user.go": `package application
type UserUseCase struct{}
type UserRepository interface{}
`,
			},
			wantIssue: "must be declared in a *_usecase.go file",
		},
		{
			name: "reports a type that does not match its file naming pattern",
			files: map[string]string{
				"internal/domain/user_repository.go": `package domain
type UserStore struct{}
`,
			},
			wantIssue: "must follow the *Repository naming pattern",
		},
		{
			name: "reports an entity outside the domain package",
			files: map[string]string{
				"internal/application/user_entity.go": `package application
type UserEntity struct{}
`,
			},
			wantIssue: "must be declared under internal/domain",
		},
		{
			name: "reports repository outside the domain repository package",
			files: map[string]string{
				"internal/infrastructure/user_repository.go": `package infrastructure
type UserRepository interface{}
`,
			},
			wantIssue: "must be declared in internal/domain/repository",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			projectDir := createProject(t, tt.files)
			issues, err := RunAnalysis(projectDir)
			if err != nil {
				t.Fatalf("RunAnalysis() error = %v", err)
			}

			if tt.wantNoIssue {
				if len(issues) != 0 {
					t.Fatalf("RunAnalysis() issues = %v; want none", issues)
				}
				return
			}

			if !containsIssue(issues, tt.wantIssue) {
				t.Fatalf("RunAnalysis() issues = %v; want one containing %q", issues, tt.wantIssue)
			}
		})
	}
}

func TestRunAnalysisMatchesConfiguredPackagesOutsideInternal(t *testing.T) {
	config := `ddd:
  layers:
    application:
      package: pkg/application
      imports_allowed: []
    infrastructure:
      package: pkg/infrastructure
      imports_allowed: []
`
	projectDir := createProjectWithConfig(t, config, map[string]string{
		"pkg/application/app.go": `package application
import "github.com/example/project/pkg/infrastructure"
var _ infrastructure.Store
`,
	})

	issues, err := RunAnalysis(projectDir)
	if err != nil {
		t.Fatalf("RunAnalysis() error = %v", err)
	}
	if !containsIssue(issues, "forbidden layer import") {
		t.Fatalf("RunAnalysis() issues = %v; want a forbidden layer import", issues)
	}
}

func TestRunAnalysisReturnsErrorsForInvalidProjects(t *testing.T) {
	t.Run("missing configuration", func(t *testing.T) {
		projectDir := t.TempDir()
		if _, err := RunAnalysis(projectDir); err == nil {
			t.Fatal("RunAnalysis() error = nil; want missing configuration error")
		}
	})

	t.Run("invalid Go source", func(t *testing.T) {
		projectDir := createProject(t, map[string]string{"internal/domain/bad.go": "package domain\nfunc {"})
		if _, err := RunAnalysis(projectDir); err == nil {
			t.Fatal("RunAnalysis() error = nil; want parse error")
		}
	})

	t.Run("nonexistent project directory", func(t *testing.T) {
		projectDir := filepath.Join(t.TempDir(), "missing")
		if _, err := RunAnalysis(projectDir); err == nil {
			t.Fatal("RunAnalysis() error = nil; want directory error")
		}
	})
}

func createProject(t *testing.T, files map[string]string) string {
	t.Helper()
	return createProjectWithConfig(t, testConfig, files)
}

func createProjectWithConfig(t *testing.T, contents string, files map[string]string) string {
	t.Helper()
	projectDir := t.TempDir()
	moduleFile := []byte("module github.com/example/project\n\ngo 1.26\n")
	if err := os.WriteFile(filepath.Join(projectDir, "go.mod"), moduleFile, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(projectDir, ".dddlint.yaml"), []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}
	for name, contents := range files {
		path := filepath.Join(projectDir, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	return projectDir
}

func containsIssue(issues []string, fragment string) bool {
	for _, issue := range issues {
		if strings.Contains(issue, fragment) {
			return true
		}
	}
	return false
}
