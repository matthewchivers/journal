package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/matthewchivers/journal/pkg/application"
	"github.com/spf13/cobra"
)

var (
	cfgPath string
	app     *application.App
)

var rootCmd = &cobra.Command{
	Use:   "journal",
	Short: "Journal is a simple CLI journaling application",
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		app = application.NewApp()
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
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "path to config file")
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
