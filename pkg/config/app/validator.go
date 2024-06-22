package app

import (
	"errors"
)

// ValidateConfig checks that the provided configuration is valid
func ValidateConfig(cfg Config) error {
	if err := validatePaths(cfg.Paths); err != nil {
		return err
	}
	if err := validateEntries(cfg.Entries, cfg.DefaultFileExtension); err != nil {
		return err
	}
	return nil
}

// validatePaths checks that the paths in the configuration are valid
func validatePaths(paths Paths) error {
	if paths.BaseDir == "" {
		return errors.New("base directory not set")
	}
	if paths.DirPattern == "" {
		return errors.New("directory pattern not set")
	}

	return nil
}

// validateEntries checks that the file types in the configuration are valid
func validateEntries(entries []Entry, defaultFileExtension string) error {
	if len(entries) == 0 {
		return errors.New("no file types defined")
	}
	for _, entry := range entries {
		if entry.ID == "" {
			return errors.New("file type name not set")
		}
		if entry.FileExtension == "" && defaultFileExtension == "" {
			return errors.New("file extension not set (defaultFileExtension must be set if fileExtension is not set for an individual file entry)")
		}
	}

	return nil
}
