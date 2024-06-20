package config

import (
	"io"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// LoadConfig loads the configuration from a file
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

	for i := range config.FileTypes {
		if config.FileTypes[i].FileExtension == "" {
			config.FileTypes[i].FileExtension = config.DefaultFileExtension
		}
	}

	return config, nil
}

// GetDefaultConfigPath returns the default path to the configuration file
func GetDefaultConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	defaultConfigPath := filepath.Join(home, ".journal", "config.yaml")
	return defaultConfigPath, nil
}
