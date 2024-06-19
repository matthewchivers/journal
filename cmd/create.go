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
	Run: func(cmd *cobra.Command, args []string) {
		templateName := cfg.DefaultFileType
		if docType != "" {
			templateName = docType
		}
		fmt.Printf("Creating new journal entry using template: %s\n", templateName)

		fileInfo := config.FileType{}
		for _, fileType := range cfg.FileTypes {
			if fileType.Name == templateName {
				fileInfo = fileType
			}
		}
		fullPath, err := paths.ConstructFullPath(cfg.Paths, fileInfo)
		if err != nil {
			fmt.Println("Error constructing file path:", err)
			os.Exit(1)
		}
		fullPath = strings.TrimSuffix(fullPath, "/")

		fileName, err := fileops.GetFileName(fileInfo)
		if err != nil {
			fmt.Println("Error getting file name:", err)
			os.Exit(1)
		}

		// Append the file name to the full path
		fullPath = fmt.Sprintf("%s/%s", fullPath, fileName)

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
