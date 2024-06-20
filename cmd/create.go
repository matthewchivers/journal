package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/matthewchivers/journal/pkg/config"
	"github.com/matthewchivers/journal/pkg/fileops"
	"github.com/matthewchivers/journal/pkg/paths"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new journal entry",
	Run: func(_ *cobra.Command, _ []string) {
		templateName := strings.ToLower(cfg.DefaultFileType)
		if docType != "" {
			templateName = strings.ToLower(docType)
		}
		fmt.Printf("Creating new journal entry using template: %s\n", templateName)

		fileInfo := config.FileType{}
		for _, fileType := range cfg.FileTypes {
			if strings.ToLower(fileType.Name) == templateName {
				fileInfo = fileType
			}
		}
		fullPath, err := paths.ConstructFullPath(cfg.Paths, fileInfo)
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
