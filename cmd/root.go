package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
	// Add subcommands here
}
