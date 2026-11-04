// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// ConfigFileName is the configuration filename in an analyzed project root.
const ConfigFileName = ".dddlint.yaml"

// Layer describes a package path and the imports permitted from that layer.
type Layer struct {
	Package        string   `yaml:"package"`
	ImportsAllowed []string `yaml:"imports_allowed"`
}

// DDDConfig contains the configured architecture layers.
type DDDConfig struct {
	Layers map[string]Layer `yaml:"layers"`
}

// Config contains the analyzer configuration loaded from a project.
type Config struct {
	ForbiddenPackagesInDomain []string   `yaml:"forbidden_packages_in_domain"`
	DDD                       DDDConfig  `yaml:"ddd"`
	Rules                     RuleConfig `yaml:"rules"`
}

// RuleConfig controls the optional naming and type-location conventions.
// Every rule defaults to enabled when omitted to preserve the original behavior.
type RuleConfig struct {
	Naming    NamingRuleConfig   `yaml:"naming"`
	Locations LocationRuleConfig `yaml:"locations"`
}

// NamingRuleConfig controls type naming conventions.
type NamingRuleConfig struct {
	Repository bool `yaml:"repository"`
	Service    bool `yaml:"service"`
	UseCase    bool `yaml:"use_case"`
}

// LocationRuleConfig controls where selected types must be declared.
type LocationRuleConfig struct {
	EntityInDomain                  bool `yaml:"entity_in_domain"`
	RepositoryInterfaceInDomainRepo bool `yaml:"repository_interface_in_domain_repository"`
	UseCaseInUseCaseFile            bool `yaml:"use_case_in_usecase_file"`
}

// LoadConfig reads and validates the analyzer configuration from a project root.
func LoadConfig(projectDir string) (*Config, error) {
	configPath := filepath.Join(projectDir, ConfigFileName)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("read configuration %s: %w", configPath, err)
	}

	cfg := Config{
		Rules: RuleConfig{
			Naming: NamingRuleConfig{
				Repository: true,
				Service:    true,
				UseCase:    true,
			},
			Locations: LocationRuleConfig{
				EntityInDomain:                  true,
				RepositoryInterfaceInDomainRepo: true,
				UseCaseInUseCaseFile:            true,
			},
		},
	}
	if err := yaml.UnmarshalStrict(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse configuration %s: %w", configPath, err)
	}
	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("validate configuration %s: %w", configPath, err)
	}

	return &cfg, nil
}

func (cfg *Config) validate() error {
	if len(cfg.DDD.Layers) == 0 {
		return fmt.Errorf("ddd.layers must define at least one layer")
	}

	seenPackages := make(map[string]string, len(cfg.DDD.Layers))
	for name, layer := range cfg.DDD.Layers {
		if strings.TrimSpace(name) == "" {
			return fmt.Errorf("layer name must not be empty")
		}
		packagePath := cleanPackagePath(layer.Package)
		if packagePath == "" || packagePath == "." || filepath.IsAbs(layer.Package) || packagePath == ".." || strings.HasPrefix(packagePath, "../") {
			return fmt.Errorf("layer %q must have a project-relative package path", name)
		}
		if previous, exists := seenPackages[packagePath]; exists {
			return fmt.Errorf("layers %q and %q use the same package path %q", previous, name, packagePath)
		}
		seenPackages[packagePath] = name

		for _, allowed := range layer.ImportsAllowed {
			allowedPath := cleanPackagePath(allowed)
			if allowedPath == "" || allowedPath == "." || filepath.IsAbs(allowed) || allowedPath == ".." || strings.HasPrefix(allowedPath, "../") {
				return fmt.Errorf("layer %q has invalid imports_allowed path %q", name, allowed)
			}
		}
	}

	return nil
}

func cleanPackagePath(value string) string {
	return strings.Trim(filepath.ToSlash(filepath.Clean(strings.TrimSpace(value))), "/")
}
