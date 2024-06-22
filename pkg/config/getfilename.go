package config

import (
	"fmt"

	"github.com/matthewchivers/journal/pkg/templating"
)

// GetFileName returns the file name for the file type
func (cfg *Config) GetFileName(entryID string) (string, error) {

	entry, err := cfg.GetEntry(entryID)
	if err != nil {
		return "", fmt.Errorf("failed to get entry: %w", err)
	}

	if entry.FileNamePattern == "" {
		return fmt.Sprintf("%s.%s", entry.ID, entry.FileExtension), nil
	}
	fileNameRaw := entry.FileNamePattern
	if fileNameRaw == "" {
		fileNameRaw = entry.ID
	}

	fileNameParsed, err := templating.ParsePattern(fileNameRaw, entry.ID, entry.FileExtension)
	if err != nil {
		return "", err
	}

	return fileNameParsed, nil
}
