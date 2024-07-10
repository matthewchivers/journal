package cmd

import (
	"os"

	"github.com/matthewchivers/journal/pkg/fileops"
	"github.com/matthewchivers/journal/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

type cliParameters struct {
	entryID       string
	directoryPath string
	baseDirectory string
	fileExtension string
	fileName      string
	topic         string
	editor        string
}

type cliFlags struct {
	noOpen bool
}

var (
	params cliParameters
	flags  cliFlags
)

var createCmd = &cobra.Command{
	Use:    "create",
	Short:  "create a new journal entry",
	PreRun: createPreRun,
	Run:    createRun,
}

func init() {
	createCmd.PersistentFlags().StringVar(&params.entryID, "id", "", "entry ID to use for templating")
	createCmd.PersistentFlags().StringVar(&params.directoryPath, "directory", "", "directory to create the file in")
	createCmd.PersistentFlags().StringVar(&params.baseDirectory, "base", "", "base directory to use")
	createCmd.PersistentFlags().StringVar(&params.fileExtension, "extension", "", "file extension to use")
	createCmd.PersistentFlags().StringVar(&params.fileName, "filename", "", "file name to use")
	createCmd.PersistentFlags().StringVar(&params.topic, "topic", "", "topic to use for templating")
	createCmd.PersistentFlags().StringVar(&params.editor, "editor", "", "editor to use for editing the file")
	createCmd.PersistentFlags().BoolVar(&flags.noOpen, "no-open", false, "do not open the file in the editor after creation")
	rootCmd.AddCommand(createCmd)
}

// createRun is the run function for the create command
// It creates a new journal entry file
func createRun(_ *cobra.Command, _ []string) {
	filePath, err := app.GetFilePath()
	if err != nil {
		logger.Log.Err(err).Msg("error getting file path")
		os.Exit(1)
	}
	logger.Log.Info().Str("file_path", filePath).
		Str("entry_id", app.EntryID).
		Msg("creating new journal entry")
	if err := fileops.CreateNewFile(filePath); err != nil {
		logger.Log.Err(err).Msg("error creating file")
		os.Exit(1)
	}
	editor, err := app.GetEditor()
	if err != nil {
		logger.Log.Err(err).Msg("error getting editor")
		os.Exit(1)
	}
	if !flags.noOpen {
		if err := editor.OpenFile(filePath); err != nil {
			logger.Log.Err(err).Msg("error opening file in editor")
			os.Exit(1)
		}
	}

}

// createPreRun is the pre-run function for the create command
// It logs the parameters and prepares the pattern/template data
func createPreRun(_ *cobra.Command, _ []string) {
	logger.Log.Debug().Dict("parameters",
		zerolog.Dict().
			Str("entry_id", params.entryID).
			Str("directory_path", params.directoryPath).
			Str("file_extension", params.fileExtension).
			Str("file_name", params.fileName).
			Str("topic", params.topic).
			Str("editor", params.editor),
	).Dict("flags",
		zerolog.Dict().Bool("no_open", flags.noOpen),
	).
		Str("command", "create").
		Msg("creating new journal entry with the 'create' command")

	if err := app.PreparePatternData(); err != nil {
		logger.Log.Err(err).Msg("error preparing pattern data")
		os.Exit(1)
	}

	if err := initialiseAppValues(); err != nil {
		logger.Log.Err(err).Msg("error setting template dependencies")
		os.Exit(1)
	}
}

// initialiseAppValues sets the values for the template dependencies
func initialiseAppValues() error {
	if err := app.SetEntryID(params.entryID); err != nil {
		return err
	}
	if err := app.SetTopic(params.topic); err != nil {
		return err
	}
	if err := app.SetFileExtension(params.fileExtension); err != nil {
		return err
	}
	if err := app.SetBaseDirectory(params.baseDirectory); err != nil {
		return err
	}

	// FileName and EntryDirectory depend on other values being set - call them last
	if err := app.SetFileName(params.fileName); err != nil {
		return err
	}
	if err := app.SetEntryDirectory(params.directoryPath); err != nil {
		return err
	}

	// Editor relies on all paths being set
	if err := app.SetEditor(params.editor); err != nil {
		return err
	}
	return nil
}
