package application

import (
	"path/filepath"

	"github.com/matthewchivers/journal/pkg/config"
	"github.com/matthewchivers/journal/pkg/logger"
	"github.com/matthewchivers/journal/pkg/paths"
)

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
