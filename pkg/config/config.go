package config

import (
	"fmt"

	"github.com/rs/zerolog"
)

var (
	logger *zerolog.Logger
)

// Config contains the configuration for the application
type Config struct {
	// DefaultEntry: specify the entry id of the desired default entry
	DefaultEntry string `yaml:"defaultEntry"`

	// DefaultFileExt is the default file extension to use when creating a new entry
	DefaultFileExt string `yaml:"defaultFileExt,omitempty"`

	// Entries is a list of entries
	Entries []Entry `yaml:"entries"`

	// Paths contains the paths to directories used by the application
	Paths Paths `yaml:"paths"`

	// UserSettings contains user-specific settings
	UserSettings UserSettings `yaml:"userSettings,omitempty"`
}

// Entry contains the configuration for an entry
func (cfg *Config) GetEntry(entryID string) (*Entry, error) {
	for _, entry := range cfg.Entries {
		if entry.ID == entryID {
			return &entry, nil
		}
	}
	return nil, fmt.Errorf("entry not found: %s", entryID)
}

// SetLogger sets the logger
func SetLogger(l *zerolog.Logger) {
	logger = l
}
