package editor

// VSCode is an editor implementation for Visual Studio Code

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/matthewchivers/journal/pkg/logger"
	"github.com/rs/zerolog"
)

var log *zerolog.Logger

type VSCode struct {
	parentDirectory string
}

// NewVSCodeEditor creates a new Visual Studio Code editor
func NewVSCodeEditor() *VSCode {
	lgr, err := logger.GetLogger()
	if err != nil {
		log = lgr
	}
	return &VSCode{}
}

// SetParentDirectory sets the parent directory for the editor
func (v *VSCode) SetParentDirectory(parentDirectory string) {
	v.parentDirectory = parentDirectory
}

// OpenFile opens a file in Visual Studio Code
func (v *VSCode) OpenFile(filePath string) error {
	log.Info().Str("file_path", filePath).
		Str("editor", "Visual Studio Code").
		Str("parent_directory", v.parentDirectory).
		Msg("opening file in Visual Studio Code")
	if v.parentDirectory == "" {
		return errors.New("parent directory not set")
	}

	// *** Security - Path Traversal ***
	// Validate all paths to ensure they are safe and do not contain any malicious content
	if err := validatePath(v.parentDirectory); err != nil {
		log.Err(err).Str("parent_directory", v.parentDirectory).
			Msg("error validating parent directory")
		return err
	}
	if err := isDirectory(v.parentDirectory); err != nil {
		log.Err(err).Str("parent_directory", v.parentDirectory).
			Msg("error validating parent directory")
		return err
	}
	if err := validatePath(filePath); err != nil {
		log.Err(err).Str("file_path", filePath).
			Msg("error validating file path")
		return err
	}

	// #nosec G204: Subprocess launched with a potential tainted input or cmd arguments
	// The inputs have been validated
	cmd := exec.Command("code", v.parentDirectory, "--goto", filePath)
	if err := cmd.Run(); err != nil {
		log.Err(err).Str("file_path", filePath).
			Str("editor", "Visual Studio Code").
			Str("parent_directory", v.parentDirectory).
			Msg("error opening file in Visual Studio Code")
		return err
	}
	log.Debug().Str("file_path", filePath).
		Str("editor", "Visual Studio Code").
		Str("parent_directory", v.parentDirectory).
		Msg("file opened in Visual Studio Code")
	return nil
}

// isDirectory validates the parent directory
func isDirectory(dir string) error {
	// check if the directory is a directory
	if fileInfo, err := os.Stat(dir); err != nil || !fileInfo.IsDir() {
		return errors.New("parent directory is not a directory")
	}
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
