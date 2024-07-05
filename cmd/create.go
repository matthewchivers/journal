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
	FileExtParam       string
	fileNameParam      string
	topicParam         string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new journal entry",
	Run: func(_ *cobra.Command, _ []string) {
		if err := app.PreparePatternData(); err != nil {
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

		filePath, err := app.GetFilePath()
		if err != nil {
			fmt.Println("error getting file path:", err)
			os.Exit(1)
		}
		fmt.Printf("Creating new journal entry using entry id: %s\n", app.EntryID)
		if err := fileops.CreateNewFile(filePath); err != nil {
			fmt.Println("error creating file:", err)
			os.Exit(1)
		}
	},
}

func init() {
	createCmd.PersistentFlags().StringVar(&entryIDParam, "id", "", "entry ID to use for templating")
	createCmd.PersistentFlags().StringVar(&directoryPathParam, "directory", "", "directory to create the file in")
	createCmd.PersistentFlags().StringVar(&FileExtParam, "extension", "", "file extension to use")
	createCmd.PersistentFlags().StringVar(&fileNameParam, "filename", "", "file name to use")
	createCmd.PersistentFlags().StringVar(&topicParam, "topic", "", "topic to use for templating")
	rootCmd.AddCommand(createCmd)
}

// setTemplateDependencies sets the values for the template dependencies
func setTemplateDependencies() error {
	if err := app.SetEntryID(entryIDParam); err != nil {
		return err
	}
	if err := app.SetTopic(topicParam); err != nil {
		return err
	}
	if err := app.SetFileExt(FileExtParam); err != nil {
		return err
	}
	return nil
}

// setTemplatedValues sets the values for the templated values
func setTemplatedValues() error {
	if err := app.SetFileName(fileNameParam); err != nil {
		return err
	}
	if err := app.SetDirectory(directoryPathParam); err != nil {
		return err
	}
	return nil
}
