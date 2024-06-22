package fileops

import (
	"os"
	"testing"

	"github.com/matthewchivers/journal/pkg/config"
	"github.com/stretchr/testify/assert"
)

// GetTempDir creates a directory for testing purposes
func GetTempDir() string {
	tempDir, err := os.MkdirTemp("", "journal")
	if err != nil {
		panic(err)
	}
	return tempDir
}

// TestCreateNewFile tests the CreateNewFile function for success and failure scenarios
func TestCreateNewFile(t *testing.T) {
	tempdir := GetTempDir()

	tests := []struct {
		name             string
		cfg              *config.Config
		expectedError    bool
		expectedErrorMsg string
	}{
		{
			name: "successful file creation",
			cfg: &config.Config{
				Paths: config.Paths{
					BaseDir:    tempdir,
					DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
				},
				Entries: []config.Entry{
					{
						ID:              "foo",
						FileExtension:   "md",
						FileNamePattern: "{{.EntryID}}.{{.FileExtension}}",
					},
				},
			},
			expectedError: false,
		},
		{
			name: "successful file creation - hardcoded extension", // support this as valid functionality in the future
			cfg: &config.Config{
				Paths: config.Paths{
					BaseDir:    tempdir,
					DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
				},
				Entries: []config.Entry{
					{
						ID:              "foo",
						FileNamePattern: "{{.EntryID}}.md",
					},
				},
			},
			expectedError: true,
		},
		{
			name: "successful file creation - no extension",
			cfg: &config.Config{
				Paths: config.Paths{
					BaseDir:    tempdir,
					DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
				},
				Entries: []config.Entry{
					{
						ID:              "foo",
						FileNamePattern: "{{.EntryID}}",
					},
				},
			},
			expectedError: false,
		},
		{
			name: "successful file creation - custom subdirectory",
			cfg: &config.Config{
				Paths: config.Paths{
					BaseDir:    tempdir,
					DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
				},
				Entries: []config.Entry{
					{
						ID:              "foo",
						FileNamePattern: "{{.EntryID}}.{{.FileExtension}}",
						SubDirPattern:   "{{.EntryID}}s",
						FileExtension:   "md",
					},
				},
			},
			expectedError: false,
		},
	}

	defer func() {
		// Clean up the file system
		os.RemoveAll(tempdir)
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry := tt.cfg.Entries[0]

			fullPath, err := tt.cfg.GetEntryPath(entry.ID)
			assert.NoError(t, err)

			// Main function under test
			err = CreateNewFile(fullPath)
			// Assert error handling
			if tt.expectedError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMsg)
			} else {
				assert.NoError(t, err)
				_, err := os.Stat(fullPath)
				assert.NoError(t, err, "file should have been created successfully")
			}
		})
	}
}
