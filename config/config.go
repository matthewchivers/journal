package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	JournalDir string     `yaml:"journalDir"`
	Documents  []Document `yaml:"documents"`
}

type Document struct {
	Name     string `yaml:"name"`
	Template string `yaml:"template"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	yamlData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(yamlData, config); err != nil {
		return nil, err
	}

	return config, nil
}
