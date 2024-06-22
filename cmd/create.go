package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/matthewchivers/journal/pkg/fileops"
	"github.com/spf13/cobra"
)

var topic string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new journal entry",
	Run: func(_ *cobra.Command, _ []string) {
		templateName := strings.ToLower(cfg.DefaultEntry)
		if docType != "" {
			templateName = strings.ToLower(docType)
		}
		fmt.Printf("Creating new journal entry using template: %s\n", templateName)

		entryPath, err := cfg.GetEntryPath(templateName)
		if err != nil {
			fmt.Println("Error getting entry path:", err)
			os.Exit(1)
		}

		if err := fileops.CreateNewFile(entryPath); err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
	},
}

func init() {
	createCmd.PersistentFlags().StringVarP(&docType, "template", "t", "", "document template to use")
	createCmd.PersistentFlags().StringVarP(&topic, "topic", "p", "", "topic to use for templating")
	rootCmd.AddCommand(createCmd)
}
