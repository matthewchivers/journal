package config

import (
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config contains the configuration for the application
type Config struct {
	DefaultFileType string       `yaml:"defaultFileType"`
	FileTypes       []FileType   `yaml:"fileTypes,omitempty"`
	Paths           Paths        `yaml:"paths"`
	UserSettings    UserSettings `yaml:"userSettings,omitempty"`
	DefaultFileExt  string       `yaml:"defaultFileExt,omitempty"`
}

// FileType contains the configuration for a file type
type FileType struct {
	Name             string   `yaml:"name"`
	FileExtension    string   `yaml:"fileExtension,omitempty"`
	Schedule         Schedule `yaml:"schedule,omitempty"`
	SubDirPattern    string   `yaml:"SubDirPattern,omitempty"`
	FileNamePattern  string   `yaml:"fileNamePattern,omitempty"`
	CustomDirPattern string   `yaml:"customDirPattern,omitempty"`
}

// Schedule contains the schedule for a file type
type Schedule struct {
	Frequency string `yaml:"frequency"`
	Interval  int    `yaml:"interval,omitempty"`
	Days      []int  `yaml:"days,omitempty"`
	Dates     []int  `yaml:"dates,omitempty"`
	Weeks     []int  `yaml:"weeks,omitempty"`
	Months    []int  `yaml:"months,omitempty"`
}

// Paths contains the paths to directories used by the application
type Paths struct {
	TemplatesDir string `yaml:"templatesDir,omitempty"`
	BaseDir      string `yaml:"baseDir"`
	DirPattern   string `yaml:"dirPattern,omitempty"`
}

// UserSettings contains user-specific settings
type UserSettings struct {
	Timezone string `yaml:"timezone,omitempty"`
}

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
