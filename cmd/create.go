package cmd

import (
	"os"

	"github.com/matthewchivers/journal/pkg/fileops"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

type parameters struct {
	entryID       string
	directoryPath string
	fileExt       string
	fileName      string
	topic         string
}

var (
	params parameters
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new journal entry",
	Run: func(_ *cobra.Command, _ []string) {
		log.Debug().Dict("parameters",
			zerolog.Dict().
				Str("entry_id", params.entryID).
				Str("directory", params.directoryPath).
				Str("extension", params.fileExt).
				Str("filename", params.fileName).
				Str("topic", params.topic),
		).
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
	createCmd.PersistentFlags().StringVar(&params.entryID, "id", "", "entry ID to use for templating")
	createCmd.PersistentFlags().StringVar(&params.directoryPath, "directory", "", "directory to create the file in")
	createCmd.PersistentFlags().StringVar(&params.fileExt, "extension", "", "file extension to use")
	createCmd.PersistentFlags().StringVar(&params.fileName, "filename", "", "file name to use")
	createCmd.PersistentFlags().StringVar(&params.topic, "topic", "", "topic to use for templating")
	rootCmd.AddCommand(createCmd)
}

// setTemplateDependencies sets the values for the template dependencies
func setTemplateDependencies() error {
	if err := app.SetEntryID(params.entryID); err != nil {
		return err
	}
	if err := app.SetTopic(params.topic); err != nil {
		return err
	}
	if err := app.SetFileExt(params.fileExt); err != nil {
		return err
	}
	return nil
}

// setTemplatedValues sets the values for the templated values
func setTemplatedValues() error {
	if err := app.SetFileName(params.fileName); err != nil {
		return err
	}
	if err := app.SetDirectory(params.directoryPath); err != nil {
		return err
	}
	return nil
}
