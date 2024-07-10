package config

import (
	"errors"
)

// Validate checks that the provided configuration is valid
func (cfg *Config) Validate() error {
	if err := validatePaths(cfg.Paths); err != nil {
		return err
	}
	if err := validateEntries(cfg.Entries, cfg.FileExtension); err != nil {
		return err
	}
	return nil
}

// validatePaths checks that the paths in the configuration are valid
func validatePaths(paths Paths) error {
	if paths.BaseDirectory == "" {
		return errors.New("base directory not set")
	}
	return nil
}

// validateEntries checks that the file types in the configuration are valid
func validateEntries(entries []Entry, fileExt string) error {
	if len(entries) == 0 {
		return errors.New("no file types defined")
	}
	for _, entry := range entries {
		if entry.ID == "" {
			return errors.New("file type name not set")
		}
		if entry.FileExtension == "" && fileExt == "" {
			return errors.New("file extension not set")
		}
	}

	return nil
}
