package config

import (
	"fmt"
)

// Config contains the configuration for the application (user settings, paths, entry types, etc.)
// The configuration is not intended to be modified during the application's lifecycle after the yaml
// file has been loaded in (values that may change should be stored in the App struct)
type Config struct {
	// DefaultEntry: specify the entry id of the desired default entry
	DefaultEntry string `yaml:"defaultEntry"`

	// Editor is the editor to use when opening files
	Editor string `yaml:"editor,omitempty"`

	// FileExtension is the file extension to use when creating a new entry (can be overridden per entry)
	FileExtension string `yaml:"fileExtension,omitempty"`

	// Entries is a list of entries
	Entries []Entry `yaml:"entries"`

	// Paths contains the paths to directories used by the application
	Paths Paths `yaml:"paths"`

	// UserSettings contains user-specific settings
	UserSettings UserSettings `yaml:"userSettings,omitempty"`
}

// NewConfig creates and returns a new Config object
func NewConfig() (*Config, error) {
	return &Config{}, nil
}

// FetchEntryByID retrieves an entry by its ID
func (cfg *Config) FetchEntryByID(entryID string) (*Entry, error) {
	for _, entry := range cfg.Entries {
		if entry.ID == entryID {
			return &entry, nil
		}
	}
	return nil, fmt.Errorf("entry not found: %s", entryID)
}
