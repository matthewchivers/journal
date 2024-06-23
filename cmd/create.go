package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/matthewchivers/journal/pkg/fileops"
	"github.com/spf13/cobra"
)

var (
	entryIDParam string
	topicParam   string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new journal entry",
	Run: func(_ *cobra.Command, _ []string) {
		entryID := GetEntryID()
		SetEntryOverrides(entryID)

		entryPath, err := cfg.GetEntryPath(entryID)
		if err != nil {
			fmt.Println("error getting entry path:", err)
			os.Exit(1)
		}

		fmt.Printf("Creating new journal entry using template: %s\n", entryID)
		if err := fileops.CreateNewFile(entryPath); err != nil {
			fmt.Println("error creating file:", err)
			os.Exit(1)
		}
	},
}

func init() {
	createCmd.PersistentFlags().StringVarP(&entryIDParam, "entry", "e", "", "document template to use")
	createCmd.PersistentFlags().StringVarP(&topicParam, "topic", "t", "", "topic to use for templating")
	rootCmd.AddCommand(createCmd)
}

func GetEntryID() string {
	entryID := strings.ToLower(cfg.DefaultEntry)
	if entryIDParam != "" {
		entryID = strings.ToLower(entryIDParam)
	}

	if entryID == "" {
		fmt.Println("no entry specified")
		os.Exit(1)
	}
	if _, err := cfg.GetEntry(entryID); err != nil {
		fmt.Println("error getting entry:", err)
		os.Exit(1)
	}
	return entryID
}

func SetEntryOverrides(entryID string) {
	if topicParam != "" {
		cfg.SetTopicForEntryID(entryID, topicParam)
	}
}
