package application

import (
	"time"

	"github.com/matthewchivers/journal/pkg/config"
	"github.com/matthewchivers/journal/pkg/editor"
	"github.com/matthewchivers/journal/pkg/logger"
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
