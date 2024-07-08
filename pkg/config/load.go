package config

import (
	"io"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// LoadConfig loads the configuration from the specified file
func (cfg *Config) LoadConfig(configPath string) error {
	log.Debug().Str("config_path", configPath).Msg("loading configuration")

	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	yamlData, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	log.Debug().Str("yaml_data", string(yamlData)).Msg("loaded yaml data")

	if err := yaml.Unmarshal(yamlData, cfg); err != nil {
		return err
	}

	return nil
}
