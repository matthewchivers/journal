package cmd

import (
	"fmt"
	"os"

	"github.com/matthewchivers/journal/pkg/config"
	"github.com/spf13/cobra"
)

var (
	cfgPath string
	cfg     *config.Config
	docType string
)

var rootCmd = &cobra.Command{
	Use:   "journal",
	Short: "Journal is a simple CLI journaling application",
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		rawCfg, err := config.LoadConfig(cfgPath)
		if err != nil {
			fmt.Println("Unable to load config file", err)
			os.Exit(1)
		}
		err = rawCfg.Validate()
		if err != nil {
			fmt.Println("Invalid config file:", err)
			os.Exit(1)
		}
		cfg = rawCfg
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
	defaultConfigPath, err := config.GetDefaultConfigPath()
	if err != nil {
		fmt.Println("Error getting default config path:", err)
		os.Exit(1)
	}

	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", defaultConfigPath, "path to config file")
}
