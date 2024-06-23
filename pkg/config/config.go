package config

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/matthewchivers/journal/pkg/templating"
)

// Config contains the configuration for the application
type Config struct {
	// launchTime is the time the application was launched
	launchTime time.Time

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

	templateModel, err := templating.PrepareTemplateData(cfg.launchTime)
	if err != nil {
		return "", fmt.Errorf("failed to prepare template data: %w", err)
	}
	templateModel.EntryID = entry.ID
	templateModel.FileExtension = entry.FileExtension

	if strings.Contains(journalDirPattern, "{{.WeekCommencing}}") {
		templateModel.AdjustForWeekCommencing(cfg.launchTime)
	}

	journalPath, err := templateModel.ParsePattern(journalDirPattern)
	if err != nil {
		return "", fmt.Errorf("failed to construct journal path: %w", err)
	}

	nestedPath, err := templateModel.ParsePattern(entry.Directory)
	if err != nil {
		return "", fmt.Errorf("failed to construct nested path: %w", err)
	}

	fileName, err := templateModel.ParsePattern(entry.FileName)
	if err != nil {
		return "", fmt.Errorf("failed to construct file name: %w", err)
	}

	fullPath := filepath.Join(cfg.Paths.BaseDirectory, journalPath, nestedPath, fileName)

	return fullPath, nil
}
