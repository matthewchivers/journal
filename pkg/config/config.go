package config

import (
	"fmt"
	"path/filepath"

	"github.com/matthewchivers/journal/pkg/templating"
)

// Config contains the configuration for the application
type Config struct {
	// DefaultEntry: specify the entry id of the desired default entry
	DefaultEntry string `yaml:"defaultEntry"`

	// DefaultFileExtension is the default file extension to use when creating a new entry
	DefaultFileExtension string `yaml:"defaultFileExtension,omitempty"`

	// Entries is a list of entries
	Entries []Entry `yaml:"entries"`

	// Paths contains the paths to directories used by the application
	Paths Paths `yaml:"paths"`

	// UserSettings contains user-specific settings
	UserSettings UserSettings `yaml:"userSettings,omitempty"`
}

func (cfg *Config) GetEntry(entryID string) (*Entry, error) {
	for _, entry := range cfg.Entries {
		if entry.ID == entryID {
			return &entry, nil
		}
	}
	return nil, fmt.Errorf("entry not found: %s", entryID)
}

// GetEntryPath returns the directory path for the entry
func (cfg *Config) GetEntryPath(entryID string) (string, error) {
	entry, err := cfg.GetEntry(entryID)
	if err != nil {
		return "", fmt.Errorf("failed to get entry: %w", err)
	}

	journalDirPattern := cfg.Paths.JournalDirectory
	if entry.JournalDirOverride != "" {
		journalDirPattern = entry.JournalDirOverride
	}

	journalPath, err := templating.ParsePattern(journalDirPattern, entry.ID, entry.FileExtension)
	if err != nil {
		return "", fmt.Errorf("failed to construct journal path: %w", err)
	}

	nestedPath, err := templating.ParsePattern(entry.Directory, entry.ID, entry.FileExtension)
	if err != nil {
		return "", fmt.Errorf("failed to construct nested path: %w", err)
	}

	fileName, err := templating.ParsePattern(entry.FileName, entry.ID, entry.FileExtension)
	if err != nil {
		return "", fmt.Errorf("failed to construct file name: %w", err)
	}

	fullPath := filepath.Join(cfg.Paths.BaseDirectory, journalPath, nestedPath, fileName)

	return fullPath, nil
}
