package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/matthewchivers/journal/config"
	"github.com/matthewchivers/journal/fileops"
	"github.com/spf13/cobra"
)

var (
	cfg     = &config.Config{}
	docType string
)

var rootCmd = &cobra.Command{
	Use:   "journal",
	Short: "Journal is a simple CLI journaling application",
	Run: func(cmd *cobra.Command, args []string) {
		templateName := cfg.DefaultDocType
		if docType != "" {
			templateName = docType
		}
		fileops.CreateNewFile(cfg, templateName)
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

	defaultConfigPath := filepath.Join(home, ".journal", "config.yaml")

	rootCmd.PersistentFlags().StringP("config", "c", defaultConfigPath, "path to config file (default: $HOME/.journal.yaml)")

	if config, err := config.LoadConfig(defaultConfigPath); err != nil {
		fmt.Println("Unable to load config file", err)
		os.Exit(1)
	} else {
		cfg = config
	}

	rootCmd.PersistentFlags().StringP("template", "t", docType, "document template to use (default: defaultDocType)")
}
