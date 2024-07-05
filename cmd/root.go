package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/matthewchivers/journal/pkg/application"
	"github.com/matthewchivers/journal/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	cfgPath          string
	loggingPathParam string
	logLevelInfo     bool
	logLevelDebug    bool
	app              *application.App
	log              *zerolog.Logger
)

var rootCmd = &cobra.Command{
	Use:   "journal",
	Short: "Journal is a simple CLI journaling application",
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		err := setupLogging()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		app = application.NewApp()
		log.Info().Msg("Starting Journal CLI")
		app.SetLaunchTime(time.Now())
		if err := loadConfig(); err != nil {
			fmt.Println("error loading config:", err)
			os.Exit(1)
		}
	},
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Welcome to Journal CLI. Use 'journal --help' to see available commands.")
	},
}

// Execute runs the root command
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgPath, "config", "", "path to config file")
	rootCmd.PersistentFlags().StringVar(&loggingPathParam, "logpath", "", "path to log file")
	rootCmd.PersistentFlags().BoolVar(&logLevelInfo, "info", false, "set log level to info")
	rootCmd.PersistentFlags().BoolVar(&logLevelDebug, "debug", false, "set log level to debug")
}

// loadConfig loads the configuration file
func loadConfig() error {
	if err := app.SetConfigPath(cfgPath); err != nil {
		return err
	}
	if err := app.SetupConfig(); err != nil {
		return err
	}
	return nil
}

// setupLogging sets the log level for the application based on the flags provided
func setupLogging() error {
	switch {
	case logLevelInfo:
		logger.SetLogLevel(logger.LogLevelInfo)
	case logLevelDebug:
		logger.SetLogLevel(logger.LogLevelDebug)
	default:
		logger.SetLogLevel(logger.LogLevelDefault)
	}

	err := logger.SetLoggingPath(loggingPathParam)
	if err != nil {
		fmt.Println("error setting logging path:", err)
		os.Exit(1)
	}

	log, err = logger.GetLogger()
	if err != nil {
		fmt.Println("error getting logger: ", err)
		return err
	}
	return nil
}
