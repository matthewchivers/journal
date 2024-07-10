package application

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/matthewchivers/journal/pkg/config"
	"github.com/matthewchivers/journal/pkg/editor"
	"github.com/matthewchivers/journal/pkg/logger"
	"github.com/matthewchivers/journal/pkg/paths"
	"github.com/matthewchivers/journal/pkg/templating"
)

type App struct {
	// LaunchTime is the time the application was launched
	LaunchTime time.Time

	// ConfigPath is the path to the configuration file
	ConfigPath string

	// Config is the application configuration
	Config *config.Config

	// EntryID is the ID of the entry
	EntryID string

	// BaseDirectory is the base directory for the entry
	BaseDirectory string

	// EntryDirectory is the directory in which to create the entry
	EntryDirectory string

	// FileName is the name of the file to create
	FileName string

	// FilePath is the full path to the file
	FilePath string

	// TemplateData is the data used to populate the templating patterns
	TemplateData *templating.TemplateModel

	// targetEntry is the targetEntry configuration (used for convenience)
	targetEntry *config.Entry

	// Editor is the ID of the editor to use
	Editor string

	// targetEditor is the editor to use
	targetEditor editor.Editor
}

// NewApp creates a new context instance
func NewApp() (*App, error) {
	app := &App{}
	return app, nil
}

// SetLaunchTime sets the launch time of the application
func (app *App) SetLaunchTime(launchTime time.Time) {
	app.LaunchTime = launchTime
	logger.Log.Debug().Str("launch_time", launchTime.String()).Msg("launch time set")
}

// SetConfigPath sets the path to the configuration file
func (app *App) SetConfigPath(cfgPath string) error {
	if cfgPath != "" {
		app.ConfigPath = cfgPath
	} else {
		appHome, err := paths.GetAppHomePath()
		if err != nil {
			return err
		}
		defaultConfigPath := filepath.Join(appHome, "config.yaml")
		app.ConfigPath = defaultConfigPath
	}
	return nil
}

// SetupConfig loads the configuration from the specified path or the default path if specified path is empty
func (app *App) SetupConfig() error {
	// Load the configuration
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	if err := cfg.LoadConfig(app.ConfigPath); err != nil {
		return err
	}
	if err := cfg.Validate(); err != nil {
		return err
	}
	app.Config = cfg
	logger.Log.Debug().Msg("config loaded and validated")
	return nil
}

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

// SetEditor sets the editor ID for the entry
// If editor is empty, the default editor is used
func (app *App) SetEditor(editorID string) error {
	if app.targetEntry == nil {
		return errors.New("entry must be set before setting editor ID")
	}
	if editorID != "" {
		app.Editor = editorID
	}
	app.Editor = app.Config.Editor
	if app.targetEntry.Editor != "" {
		app.Editor = app.targetEntry.Editor
	}

	switch app.Editor {
	case "":
		return errors.New("editor not set")
	case "vscode":
		vscEditor, err := editor.NewVSCodeEditor()
		if err != nil {
			return err
		}
		app.targetEditor = vscEditor
	default:
		return errors.New("editor not supported")
	}

	return nil
}

// GetEditor returns the editor for the entry
func (app *App) GetEditor() (editor.Editor, error) {
	if app.targetEditor == nil {
		return nil, errors.New("editor must be set before getting editor")
	}
	return app.targetEditor, nil
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

// PreparePatternData prepares the pattern data for the application
// This must be called before parsing any patterns
func (app *App) PreparePatternData() error {
	if app.LaunchTime.IsZero() {
		return errors.New("launch time must be set before preparing pattern data")
	}
	templateModel, err := templating.PrepareTemplateData(app.LaunchTime)
	if err != nil {
		return fmt.Errorf("failed to prepare template data: %w", err)
	}

	app.TemplateData = &templateModel
	return nil
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

// SetTopic sets the topic for the entry
func (app *App) SetTopic(topic string) error {
	if app.TemplateData == nil {
		return errors.New("pattern data must be initialised before setting topic")
	}
	if topic != "" {
		app.TemplateData.Topic = topic
	} else {
		entry, err := app.GetTargetEntry()
		if err != nil {
			return err
		}
		app.TemplateData.Topic = entry.Topic
	}
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
