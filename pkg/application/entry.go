package application

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/matthewchivers/journal/pkg/config"
	"github.com/matthewchivers/journal/pkg/logger"
)

// SetEntryID sets the entry ID for the context
// If entryID is empty, the default entry is used
func (app *App) SetEntryID(entryID string) error {
	if app.Config == nil {
		return errors.New("config must be loaded before setting entry ID")
	}
	if app.TemplateData == nil {
		return errors.New("pattern data must be initialised before setting entry id")
	}
	if entryID != "" {
		app.EntryID = strings.ToLower(entryID)
	} else {
		app.EntryID = strings.ToLower(app.Config.DefaultEntry)
	}
	if app.EntryID == "" {
		return errors.New("no entry specified")
	}
	if _, err := app.Config.FetchEntryByID(app.EntryID); err != nil {
		return err
	}

	app.TemplateData.EntryID = app.EntryID

	return nil
}

// GetTargetEntry returns the target entry configuration
// If the target entry is not already set, it is fetched from the configuration
func (app *App) GetTargetEntry() (*config.Entry, error) {
	if app.targetEntry != nil {
		return app.targetEntry, nil
	}
	if app.Config == nil {
		return nil, errors.New("config must be loaded before getting entry")
	}
	if app.EntryID == "" {
		return nil, errors.New("entry ID must be set before getting entry")
	}
	ent, err := app.Config.FetchEntryByID(app.EntryID)
	if err != nil {
		return nil, err
	}
	app.targetEntry = ent
	return ent, nil
}

// SetEntryDirectory sets the directory for the entry
// If dir is empty, the directory is calculated from the entry configuration
func (app *App) SetEntryDirectory(dir string) error {
	if dir == "" {
		entryPath, err := app.CalculateEntryDir()
		if err != nil {
			return err
		}
		app.EntryDirectory = entryPath
	} else {
		app.EntryDirectory = dir
	}
	return nil
}

// CalculateEntryDir calculates and returns the directory for the entry
// Does not set the directory in the context
func (app *App) CalculateEntryDir() (string, error) {
	entry, err := app.GetTargetEntry()
	if err != nil {
		return "", err
	}

	if app.TemplateData == nil {
		err := errors.New("template data must be initialised before getting entry directory")
		logger.Log.Err(err).Msg("")
		return "", err
	}

	if app.BaseDirectory == "" {
		err := errors.New("base directory must be set before getting entry directory")
		logger.Log.Err(err).Msg("")
		return "", err
	}

	entryDirectory, err := app.TemplateData.ParsePattern(entry.DirectoryPattern)
	if err != nil {
		return "", fmt.Errorf("failed to construct nested path: %w", err)
	}

	fullDirectory := filepath.Join(app.BaseDirectory, entryDirectory)

	logger.Log.Debug().Str("base_dir", app.Config.Paths.BaseDirectory).
		Str("entry_dir_pattern", entry.DirectoryPattern).
		Str("entry_dir_pattern_parsed", entryDirectory).
		Str("full_directory path", fullDirectory).
		Msg("entry directory calculation (basedir/journaldir/entrydir)")

	return fullDirectory, nil
}

// GetEntryFileName returns the file name for the entry
// If the file name is not set, the file name is calculated from the entry configuration
func (app *App) GetEntryFileName() (string, error) {
	if app.FileName != "" {
		return app.FileName, nil
	}
	entry, err := app.GetTargetEntry()
	if err != nil {
		return "", err
	}

	if app.TemplateData == nil {
		return "", errors.New("pattern data must be initialised before getting entry file name")
	}

	fileName, err := app.TemplateData.ParsePattern(entry.FileNamePattern)
	if err != nil {
		return "", fmt.Errorf("failed to construct file name: %w", err)
	}

	logger.Log.Debug().Str("entry_id", app.EntryID).
		Str("file_name_pattern", entry.FileNamePattern).
		Str("file_name_parsed", fileName).
		Msg("entry filename calculation")
	return fileName, nil
}

// SetFileExtension sets the file extension for the entry
// If fileExt is empty, the file extension is calculated from the configuration
// (entry, or default if entry does not specify an extension)
func (app *App) SetFileExtension(fileExt string) error {
	if app.TemplateData == nil {
		return errors.New("template data must be initialised before setting file extension")
	}
	if fileExt != "" {
		app.TemplateData.FileExtension = fileExt
	} else {
		entry, err := app.GetTargetEntry()
		if err != nil {
			return err
		}
		if entry.FileExtension != "" {
			app.TemplateData.FileExtension = entry.FileExtension
		} else {
			app.TemplateData.FileExtension = app.Config.FileExtension
		}
	}
	return nil
}
