// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	projectDir := t.TempDir()
	configPath := filepath.Join(projectDir, ConfigFileName)
	contents := `forbidden_packages_in_domain:
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
`
	if err := os.WriteFile(configPath, []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}

	got, err := LoadConfig(projectDir)
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}
	if got.DDD.Layers["domain"].Package != "internal/domain" {
		t.Errorf("domain package = %q; want %q", got.DDD.Layers["domain"].Package, "internal/domain")
	}
	if len(got.DDD.Layers["application"].ImportsAllowed) != 1 || got.DDD.Layers["application"].ImportsAllowed[0] != "internal/domain" {
		t.Errorf("application imports_allowed = %#v; want [internal/domain]", got.DDD.Layers["application"].ImportsAllowed)
	}
	if len(got.ForbiddenPackagesInDomain) != 1 || got.ForbiddenPackagesInDomain[0] != "net/http" {
		t.Errorf("forbidden packages = %#v; want [net/http]", got.ForbiddenPackagesInDomain)
	}
}

func TestLoadConfigErrors(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{name: "missing file", wantErr: true},
		{name: "invalid yaml", content: "ddd: [", wantErr: true},
		{name: "unknown field", content: "unexpected: true", wantErr: true},
		{name: "missing layers", content: "ddd:\n  layers: {}", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			projectDir := t.TempDir()
			if tt.content != "" {
				if err := os.WriteFile(filepath.Join(projectDir, ConfigFileName), []byte(tt.content), 0o600); err != nil {
					t.Fatal(err)
				}
			}

			_, err := LoadConfig(projectDir)
			if (err != nil) != tt.wantErr {
				t.Fatalf("LoadConfig() error = %v; want error %v", err, tt.wantErr)
			}
		})
	}
}
