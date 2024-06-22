package config

import (
	"fmt"
	"path/filepath"

	"github.com/matthewchivers/journal/pkg/templating"
)

func (cfg *Config) GetEntryPath(entryID string) (string, error) {
	entry, err := cfg.GetEntry(entryID)
	if err != nil {
		return "", fmt.Errorf("failed to get entry: %w", err)
	}

	mainPathPattern := cfg.Paths.DirPattern
	if entry.CustomDirPattern != "" {
		mainPathPattern = entry.CustomDirPattern
	}

	nestedPathPattern := ""
	if entry.SubDirPattern != "" {
		nestedPathPattern = entry.SubDirPattern
	}

	mainPath, err := templating.ParsePattern(mainPathPattern, entry.ID, entry.FileExtension)
	if err != nil {
		return "", fmt.Errorf("failed to construct base path: %w", err)
	}

	nestedPath, err := templating.ParsePattern(nestedPathPattern, entry.ID, entry.FileExtension)
	if err != nil {
		return "", fmt.Errorf("failed to construct nested path: %w", err)
	}

	fileName, err := templating.ParsePattern(entry.FileNamePattern, entry.ID, entry.FileExtension)
	if err != nil {
		return "", fmt.Errorf("failed to construct file name: %w", err)
	}

	fullPath := filepath.Join(cfg.Paths.BaseDir, mainPath, nestedPath, fileName)

	return fullPath, nil
}
