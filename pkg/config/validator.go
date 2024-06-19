package config

import (
	"errors"
)

// ValidateConfig checks that the provided configuration is valid
func ValidateConfig(cfg Config) error {
	if cfg.Paths.BaseDir == "" {
		return errors.New("base directory not set")
	}
	if cfg.Paths.DirPattern == "" {
		return errors.New("directory pattern not set")
	}
	if len(cfg.FileTypes) == 0 {
		return errors.New("no file types defined")
	}
	for _, fileType := range cfg.FileTypes {
		if fileType.Name == "" {
			return errors.New("file type name not set")
		}
		if fileType.FileExtension == "" && cfg.DefaultFileExtension == "" {
			return errors.New("file extension not set (defaultFileExtension must be set if fileExtension is not set for an individual file entry)")
		}
	}
	return nil
}
