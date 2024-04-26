package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

// Config contains the configuration for the application
type Config struct {
	DefaultDocType string         `yaml:"defaultDocType"`
	DocumentTypes  []DocumentType `yaml:"documentTypes"`
	Paths          Paths          `yaml:"paths"`
	UserSettings   UserSettings   `yaml:"userSettings,omitempty"`
}

// DocumentType contains the configuration for a document type
type DocumentType struct {
	Name               string   `yaml:"name"`
	Schedule           Schedule `yaml:"schedule,omitempty"`
	NestedPathTemplate string   `yaml:"nestedPathTemplate,omitempty"`
}

// Schedule contains the schedule for a document type
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
	TemplatesDir       string `yaml:"templatesDir,omitempty"`
	JournalBaseDir     string `yaml:"journalBaseDir"`
	NestedPathTemplate string `yaml:"nestedPathTemplate,omitempty"`
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
