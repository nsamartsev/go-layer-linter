package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type LayerRule struct {
	Layer  string `yaml:"layer"`
	Allows string `yaml:"allows"`
}

type LinterConfig struct {
	LayeredImports []LayerRule `yaml:"layered_imports"`
}

func LoadConfig(path string) *LinterConfig {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Can't read config: %v", err)
	}

	var config LinterConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Can't parse config: %v", err)
	}

	return &config
}
