package application

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"

	"github.com/matthewchivers/journal/pkg/config"
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

	// Directory is the directory in which to create the entry
	Directory string

	// FileName is the name of the file to create
	FileName string

	// FilePath is the full path to the file
	FilePath string

	// TemplateData is the data used to populate the templating patterns
	TemplateData *templating.TemplateModel

	// Logger is a zerolog logger
	logger *zerolog.Logger

	// entry is the entry configuration (used for convenience)
	entry *config.Entry
}

// NewApp creates a new context instance
func NewApp() *App {
	app := &App{}
	// setup logger
	app.SetLogger(getDefaultLogger())
	app.SetLaunchTime(time.Now())
	return &App{}
}

// SetLogger sets the logger for the application
func (app *App) SetLogger(logger *zerolog.Logger) {
	app.logger = logger
}

// GetLogger returns the logger for the application
// If no logger is set, a default logger is returned
func (app *App) GetLogger() *zerolog.Logger {
	if app.logger == nil {
		app.logger = getDefaultLogger()
	}
	return app.logger
}

// getDefaultLogger sets the default logger for the application
func getDefaultLogger() *zerolog.Logger {
	// Human readable console logger
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}

	newLogger := zerolog.New(consoleWriter.Out).With().Timestamp().Logger()
	return &newLogger
}

// SetLaunchTime sets the launch time of the application
func (app *App) SetLaunchTime(launchTime time.Time) {
	app.LaunchTime = launchTime
}

// SetConfigPath sets the path to the configuration file
func (app *App) SetConfigPath(cfgPathOverride string) error {
	if cfgPathOverride != "" {
		app.ConfigPath = cfgPathOverride
	} else {
		appHome, err := GetAppHomePath()
		if err != nil {
			return err
		}
		defaultConfigPath := filepath.Join(appHome, "config.yaml")
		app.ConfigPath = defaultConfigPath
	}
	return nil
}

// GetAppHomePath returns the path to the application
func GetAppHomePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	appHomePath := filepath.Join(home, ".journal")
	return appHomePath, nil
}

// SetupConfig loads the configuration from the specified path or the default path if specified path is empty
func (app *App) SetupConfig() error {
	// Load the configuration
	cfg, err := config.LoadConfig(app.ConfigPath)
	if err != nil {
		return err
	}
	if err := cfg.Validate(); err != nil {
		return err
	}
	app.Config = cfg
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
	if _, err := app.Config.GetEntry(app.EntryID); err != nil {
		return err
	}

	app.TemplateData.EntryID = app.EntryID
	return nil
}

func (app *App) SetDirectory(dir string) error {
	if dir == "" {
		entryPath, err := app.GetEntryDir()
		if err != nil {
			return err
		}
		app.Directory = entryPath
	} else {
		app.Directory = dir
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
	return nil
}

func (app *App) GetEntryDir() (string, error) {
	entry, err := app.GetEntry()
	if err != nil {
		return "", err
	}

	journalDirPattern := app.Config.Paths.JournalDirectory
	if entry.JournalDirOverride != "" {
		journalDirPattern = entry.JournalDirOverride
	}

	if app.TemplateData == nil {
		return "", errors.New("template data must be initialised before getting entry directory")
	}

	journalPath, err := app.TemplateData.ParsePattern(journalDirPattern)
	if err != nil {
		return "", fmt.Errorf("failed to construct journal path: %w", err)
	}

	nestedPath, err := app.TemplateData.ParsePattern(entry.Directory)
	if err != nil {
		return "", fmt.Errorf("failed to construct nested path: %w", err)
	}

	fullPath := filepath.Join(app.Config.Paths.BaseDirectory, journalPath, nestedPath)

	return fullPath, nil
}

// GetEntryFileName returns the file name for the entry
func (app *App) GetEntryFileName() (string, error) {
	entry, err := app.GetEntry()
	if err != nil {
		return "", err
	}

	if app.TemplateData == nil {
		return "", errors.New("pattern data must be initialised before getting entry file name")
	}

	fileName, err := app.TemplateData.ParsePattern(entry.FileName)
	if err != nil {
		return "", fmt.Errorf("failed to construct file name: %w", err)
	}

	return fileName, nil
}

// GetFilePath returns the full path to the file
func (app *App) GetFilePath() (string, error) {
	if app.FilePath == "" {
		if app.Directory == "" {
			return "", errors.New("directory must be set before getting file path")
		}
		if app.FileName == "" {
			return "", errors.New("file name must be set before getting file path")
		}
		app.FilePath = filepath.Join(app.Directory, app.FileName)
	}
	return app.FilePath, nil
}

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

// SetFileExt sets the file extension for the entry
func (app *App) SetFileExt(fileExt string) error {
	if app.TemplateData == nil {
		return errors.New("template data must be initialised before setting file extension")
	}
	if fileExt != "" {
		app.TemplateData.FileExt = fileExt
	} else {
		entry, err := app.Config.GetEntry(app.EntryID)
		if err != nil {
			return err
		}
		if entry.FileExt != "" {
			app.TemplateData.FileExt = entry.FileExt
		} else {
			app.TemplateData.FileExt = app.Config.DefaultFileExt
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
		entry, err := app.GetEntry()
		if err != nil {
			return err
		}
		app.TemplateData.Topic = entry.Topic
	}
	return nil
}

func (app *App) GetEntry() (*config.Entry, error) {
	if app.entry != nil {
		return app.entry, nil
	}
	if app.Config == nil {
		return nil, errors.New("config must be loaded before getting entry")
	}
	if app.EntryID == "" {
		return nil, errors.New("entry ID must be set before getting entry")
	}
	ent, err := app.Config.GetEntry(app.EntryID)
	if err != nil {
		return nil, err
	}
	app.entry = ent
	return ent, nil
}
