package editor

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/matthewchivers/journal/pkg/logger"
)

type VSCode struct{}

// NewVSCodeEditor creates a new Visual Studio Code editor
func NewVSCodeEditor() (*VSCode, error) {
	return &VSCode{}, nil
}

// OpenFile opens a file in Visual Studio Code
func (v *VSCode) OpenFile(filePath string) error {
	logger.Log.Info().Str("file_path", filePath).
		Str("editor", "Visual Studio Code").
		Msg("opening file in Visual Studio Code")

	// *** Security - Path Traversal ***
	// Validate path to ensure it is safe and does not contain any malicious content
	if err := validatePath(filePath); err != nil {
		logger.Log.Err(err).Str("file_path", filePath).
			Msg("error validating file path")
		return err
	}

	// #nosec G204: Subprocess launched with a potential tainted input or cmd arguments
	// The inputs have been validated
	cmd := exec.Command("code", filePath)
	if err := cmd.Run(); err != nil {
		logger.Log.Err(err).Str("file_path", filePath).
			Str("editor", "Visual Studio Code").
			Msg("error opening file in Visual Studio Code")
		return err
	}
	logger.Log.Info().Str("file_path", filePath).
		Str("editor", "Visual Studio Code").
		Msg("opened file in Visual Studio Code")
	return nil
}

// validatePath validates a path
func validatePath(path string) error {
	// allowlist: only allow characters/runes that are: alphanumeric, dots, slashes, hyphens, underscores, or spaces
	validPathPattern := `^[a-zA-Z0-9._/\- ]+$`
	if matched, _ := regexp.MatchString(validPathPattern, path); !matched {
		return errors.New("parent directory contains invalid characters")
	}
	// check if the directory is relative
	if !filepath.IsAbs(path) {
		return errors.New("parent directory must be absolute")
	}
	// check if the path / file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("parent directory does not exist")
	}
	return nil
}
