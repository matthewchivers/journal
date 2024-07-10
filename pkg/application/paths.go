package application

import (
	"errors"
	"path/filepath"

	"github.com/matthewchivers/journal/pkg/logger"
)

// SetBaseDirectory sets the base directory for the entry
// If baseDir is empty, the default base directory is used
func (app *App) SetBaseDirectory(baseDir string) error {
	if baseDir != "" {
		app.BaseDirectory = baseDir
		return nil
	}
	if app.targetEntry == nil {
		return errors.New("entry must be set before setting base directory")
	}
	app.BaseDirectory = app.Config.Paths.BaseDirectory
	if app.targetEntry.BaseDirectory != "" {
		app.BaseDirectory = app.targetEntry.BaseDirectory
	}
	return nil
}

// SetFileName sets the file name for the entry
// If fileName is empty, the default file name is retrieved
func (app *App) SetFileName(fileName string) error {
	if app.TemplateData == nil {
		return errors.New("pattern data must be initialised before setting file name")
	}
	if fileName != "" {
		app.FileName = fileName
	} else {
		fileName, err := app.GetEntryFileName()
		if err != nil {
			return err
		}
		app.FileName = fileName
	}
	logger.Log.Debug().Str("file_name_override", fileName).
		Str("file_name_final", app.FileName).
		Msg("file name set")
	return nil
}

// GetFilePath returns the full path to the file
// If the file path is not set, it is calculated from the directory and file name
func (app *App) GetFilePath() (string, error) {
	if app.FilePath != "" {
		return app.FilePath, nil
	}
	if app.EntryDirectory == "" {
		return "", errors.New("directory must be set before getting file path")
	}
	if app.FileName == "" {
		return "", errors.New("file name must be set before getting file path")
	}
	app.FilePath = filepath.Join(app.EntryDirectory, app.FileName)
	logger.Log.Debug().Str("directory", app.EntryDirectory).
		Str("file_name", app.FileName).
		Str("file_path", app.FilePath).
		Msg("file path calculated from directory and file name")
	return app.FilePath, nil
}
