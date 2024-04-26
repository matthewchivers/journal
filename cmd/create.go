package cmd

import (
	"fmt"
	"os"

	"github.com/matthewchivers/journal/pkg/fileops"
	"github.com/matthewchivers/journal/pkg/paths"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new journal entry",
	Run: func(cmd *cobra.Command, args []string) {
		templateName := cfg.DefaultDocType
		if docType != "" {
			templateName = docType
		}
		fmt.Printf("Creating new journal entry using template: %s\n", templateName)

		fullPath, err := paths.ConstructFullPath(cfg.Paths.JournalBaseDir, cfg.Paths.NestedPathTemplate, templateName)
		if err != nil {
			fmt.Println("Error constructing file path:", err)
			os.Exit(1)
		}
		if err := fileops.CreateNewFile(fullPath); err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
	},
}

func init() {
	createCmd.PersistentFlags().StringVarP(&docType, "template", "t", "", "document template to use")
	rootCmd.AddCommand(createCmd)
}
