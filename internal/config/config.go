package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Layer struct {
	Package        string   `yaml:"package"`
	ImportsAllowed []string `yaml:"imports_allowed"`
}

type Config struct {
	ForbiddenPackagesInDomain []string         `yaml:"forbidden_packages_in_domain"`
	Layers                    map[string]Layer `yaml:"ddd.layers"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile(".dddlint.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
