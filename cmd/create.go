package cmd

import (
	"fmt"
	"os"

	"github.com/matthewchivers/journal/pkg/fileops"
	"github.com/spf13/cobra"
)

var (
	entryIDParam       string
	directoryPathParam string
	fileExtensionParam string
	fileNameParam      string
	topicParam         string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new journal entry",
	Run: func(_ *cobra.Command, _ []string) {
		if err := appCtx.PreparePatternData(); err != nil {
			fmt.Println("error preparing pattern data:", err)
			os.Exit(1)
		}

		if err := setTemplateDependencies(); err != nil {
			fmt.Println("error setting template dependencies:", err)
			os.Exit(1)
		}

		if err := setTemplatedValues(); err != nil {
			fmt.Println("error setting templated values:", err)
			os.Exit(1)
		}

		filePath, err := appCtx.GetFilePath()
		if err != nil {
			fmt.Println("error getting file path:", err)
			os.Exit(1)
		}
		fmt.Printf("Creating new journal entry using entry id: %s\n", appCtx.EntryID)
		if err := fileops.CreateNewFile(filePath); err != nil {
			fmt.Println("error creating file:", err)
			os.Exit(1)
		}
	},
}

func init() {
	createCmd.PersistentFlags().StringVarP(&entryIDParam, "entryid", "id", "", "entry ID to use for templating")
	createCmd.PersistentFlags().StringVarP(&directoryPathParam, "directory", "d", "", "directory to create the file in")
	createCmd.PersistentFlags().StringVarP(&fileExtensionParam, "extension", "e", "", "file extension to use")
	createCmd.PersistentFlags().StringVarP(&fileNameParam, "filename", "f", "", "file name to use")
	createCmd.PersistentFlags().StringVarP(&topicParam, "topic", "t", "", "topic to use for templating")
	rootCmd.AddCommand(createCmd)
}

// setTemplateDependencies sets the values for the template dependencies
func setTemplateDependencies() error {
	if err := appCtx.SetEntryID(entryIDParam); err != nil {
		return err
	}
	if err := appCtx.SetTopic(topicParam); err != nil {
		return err
	}
	if err := appCtx.SetFileExtension(fileExtensionParam); err != nil {
		return err
	}
	return nil
}

// setTemplatedValues sets the values for the templated values
func setTemplatedValues() error {
	if err := appCtx.SetFileName(fileNameParam); err != nil {
		return err
	}
	if err := appCtx.SetDirectory(directoryPathParam); err != nil {
		return err
	}
	return nil
}
