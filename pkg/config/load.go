package config

import (
	"io"
	"os"

	yaml "gopkg.in/yaml.v2"
)

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
