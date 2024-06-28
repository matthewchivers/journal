package config

import (
	"io"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// GetDefaultConfigPath returns the default path to the configuration file
func GetDefaultConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	defaultConfigPath := filepath.Join(home, ".journal", "config.yaml")
	return defaultConfigPath, nil
}

// LoadConfig loads the configuration from a file and returns the configuration object
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

	for i := range config.Entries {
		if config.Entries[i].FileExt == "" {
			config.Entries[i].FileExt = config.DefaultFileExt
		}
	}

	return config, nil
}
