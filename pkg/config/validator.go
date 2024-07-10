package config

import (
	"errors"
)

// Validate checks that the provided configuration is valid
func (cfg *Config) Validate() error {
	if err := validatePaths(cfg.Paths); err != nil {
		return err
	}
	if err := validateEntries(cfg.Entries, cfg.FileExt); err != nil {
		return err
	}
	return nil
}

// validatePaths checks that the paths in the configuration are valid
func validatePaths(paths Paths) error {
	if paths.BaseDirectory == "" {
		return errors.New("base directory not set")
	}
	if paths.JournalDirectory == "" {
		return errors.New("directory pattern not set")
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
		if entry.FileExt == "" && fileExt == "" {
			return errors.New("file extension not set (FileExt must be set if fileExt is not set for an individual file entry)")
		}
	}

	return nil
}
