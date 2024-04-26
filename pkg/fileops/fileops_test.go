package fileops

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/matthewchivers/journal/pkg/config"
	"github.com/matthewchivers/journal/pkg/paths"
	"github.com/stretchr/testify/assert"
)

// MockConfig creates a configuration for testing purposes
func MockConfig() *config.Config {
	tempDir, err := os.MkdirTemp("", "journal")
	if err != nil {
		panic(err)
	}
	return &config.Config{
		Paths: config.Paths{
			BaseDir:    tempDir,
			DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
		},
	}
}

// TestCreateNewFile tests the CreateNewFile function for success and failure scenarios
func TestCreateNewFile(t *testing.T) {
	cfg := MockConfig()

	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")

	// Probably doesn't need to be a table test, but this works in the name of speed / getting an MVP.
	// Doesn't test failure to create a file, as that would border on testing the os package
	tests := []struct {
		name             string
		docTemplateName  string
		expectedError    bool
		expectedErrorMsg string
	}{
		{
			name:            "successful file creation",
			docTemplateName: "testDoc",
			expectedError:   false,
		},
	}

	defer func() {
		// Clean up the file system
		os.RemoveAll(cfg.Paths.BaseDir)
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute the function under test
			fmt.Printf("Creating new journal entry using template: %s\n", tt.docTemplateName)
			fullPath, pathErr := paths.ConstructFullPath(cfg.Paths.BaseDir, cfg.Paths.DirPattern, tt.docTemplateName)
			if pathErr != nil {
				panic(pathErr)
			}
			err := CreateNewFile(fullPath)

			// Assert error handling
			if tt.expectedError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMsg)
			} else {
				assert.NoError(t, err)
				expectedPath := filepath.Join(cfg.Paths.BaseDir, year, month, day, fmt.Sprintf("%s.md", tt.docTemplateName))
				stats, err := os.Stat(expectedPath)
				assert.NoError(t, err, "file should have been created successfully")
				fmt.Printf("Stats: %+v\n", stats)
			}
		})
	}
}
