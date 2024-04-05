package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/matthewchivers/journal/config"
	"github.com/spf13/cobra"
)

var cfgFile string

var cfg = &config.Config{}

var rootCmd = &cobra.Command{
	Use:   "journal",
	Short: "Journal is a simple CLI journaling application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func init() {

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Unable to determine user home directory", err)
		os.Exit(1)
	}

	defaultConfigPath := filepath.Join(home, ".journal.yaml")

	rootCmd.PersistentFlags().StringP("config", "c", defaultConfigPath, "config file (default is $HOME/.journal.yaml)")

	if config, err := config.LoadConfig(defaultConfigPath); err != nil {
		fmt.Println("Unable to load config file", err)
		os.Exit(1)
	} else {
		cfg = config
	}
}
