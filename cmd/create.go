package cmd

import (
	"os"

	"github.com/matthewchivers/journal/pkg/fileops"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	entryID       string
	directoryPath string
	fileExt       string
	fileName      string
	topic         string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new journal entry",
	Run: func(_ *cobra.Command, _ []string) {
		log.Debug().Dict("parameters", zerolog.Dict().
			Str("entry_id", entryID).
			Str("directory", directoryPath).
			Str("extension", fileExt).
			Str("filename", fileName).
			Str("topic", topic)).
			Str("command", "create").
			Msg("creating new journal entry with the 'create' command")
		if err := app.PreparePatternData(); err != nil {
			log.Err(err).Msg("error preparing pattern data")
			os.Exit(1)
		}

		if err := setTemplateDependencies(); err != nil {
			log.Err(err).Msg("error setting template dependencies")
			os.Exit(1)
		}

		if err := setTemplatedValues(); err != nil {
			log.Err(err).Msg("error setting templated values")
			os.Exit(1)
		}

		filePath, err := app.GetFilePath()
		if err != nil {
			log.Err(err).Msg("error getting file path")
			os.Exit(1)
		}
		log.Info().Str("file_path", filePath).
			Str("entry_id", app.EntryID).
			Msg("creating new journal entry")
		if err := fileops.CreateNewFile(filePath); err != nil {
			log.Err(err).Msg("error creating file")
			os.Exit(1)
		}
	},
}

func init() {
	createCmd.PersistentFlags().StringVar(&entryID, "id", "", "entry ID to use for templating")
	createCmd.PersistentFlags().StringVar(&directoryPath, "directory", "", "directory to create the file in")
	createCmd.PersistentFlags().StringVar(&fileExt, "extension", "", "file extension to use")
	createCmd.PersistentFlags().StringVar(&fileName, "filename", "", "file name to use")
	createCmd.PersistentFlags().StringVar(&topic, "topic", "", "topic to use for templating")
	rootCmd.AddCommand(createCmd)
}

// setTemplateDependencies sets the values for the template dependencies
func setTemplateDependencies() error {
	if err := app.SetEntryID(entryID); err != nil {
		return err
	}
	if err := app.SetTopic(topic); err != nil {
		return err
	}
	if err := app.SetFileExt(fileExt); err != nil {
		return err
	}
	return nil
}

// setTemplatedValues sets the values for the templated values
func setTemplatedValues() error {
	if err := app.SetFileName(fileName); err != nil {
		return err
	}
	if err := app.SetDirectory(directoryPath); err != nil {
		return err
	}
	return nil
}
