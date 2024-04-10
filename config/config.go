package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DefaultDocType      string         `yaml:"defaultDocType"`
	DocumentTypes       []DocumentType `yaml:"documentTypes"`
	Paths               Paths          `yaml:"paths"`
	UserSettings        UserSettings   `yaml:"userSettings,omitempty"`
	DocumentNestingPath string         `yaml:"documentNestingPath,omitempty"`
}

type DocumentType struct {
	Name                string   `yaml:"name"`
	Schedule            Schedule `yaml:"schedule,omitempty"`
	DocumentNestingPath string   `yaml:"documentNestingPath,omitempty"`
}

type Schedule struct {
	Frequency string `yaml:"frequency"`
	Interval  int    `yaml:"interval,omitempty"`
	Days      []int  `yaml:"days,omitempty"`
	Dates     []int  `yaml:"dates,omitempty"`
	Weeks     []int  `yaml:"weeks,omitempty"`
	Months    []int  `yaml:"months,omitempty"`
}

type Paths struct {
	TemplatesDir   string `yaml:"templatesDir,omitempty"`
	JournalBaseDir string `yaml:"journalBaseDir"`
}

type UserSettings struct {
	Timezone string `yaml:"timezone,omitempty"`
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
