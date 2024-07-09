package editor

// VSCode is an editor implementation for Visual Studio Code

import (
	"errors"
	"os/exec"

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
