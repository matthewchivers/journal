package fileops

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	app "github.com/matthewchivers/journal/pkg/application"
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
					BaseDirectory:    tempdir,
					JournalDirectory: "{{.Year}}/{{.Month}}/{{.Day}}",
				},
				Entries: []config.Entry{
					{
						ID:       "foo",
						FileExt:  "md",
						FileName: "{{.EntryID}}.{{.FileExt}}",
					},
				},
			},
			expectedError: false,
		},
		{
			name: "successful file creation - hardcoded extension", // support this as valid functionality in the future
			cfg: &config.Config{
				Paths: config.Paths{
					BaseDirectory:    tempdir,
					JournalDirectory: "{{.Year}}/{{.Month}}/{{.Day}}",
				},
				Entries: []config.Entry{
					{
						ID:       "foo",
						FileName: "{{.EntryID}}.md",
					},
				},
			},
			expectedError: true,
		},
		{
			name: "successful file creation - no extension",
			cfg: &config.Config{
				Paths: config.Paths{
					BaseDirectory:    tempdir,
					JournalDirectory: "{{.Year}}/{{.Month}}/{{.Day}}",
				},
				Entries: []config.Entry{
					{
						ID:       "foo",
						FileName: "{{.EntryID}}",
					},
				},
			},
			expectedError: false,
		},
		{
			name: "successful file creation - custom subdirectory",
			cfg: &config.Config{
				Paths: config.Paths{
					BaseDirectory:    tempdir,
					JournalDirectory: "{{.Year}}/{{.Month}}/{{.Day}}",
				},
				Entries: []config.Entry{
					{
						ID:        "foo",
						FileName:  "{{.EntryID}}.{{.FileExt}}",
						Directory: "{{.EntryID}}s",
						FileExt:   "md",
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

			appCtx := &app.App{
				Config: tt.cfg,
			}

			appCtx.SetLaunchTime(time.Now())

			err := appCtx.PreparePatternData()
			assert.NoError(t, err)

			err = appCtx.SetEntryID(entry.ID)
			assert.NoError(t, err)

			err = appCtx.SetFileName("")
			assert.NoError(t, err)

			err = appCtx.SetDirectory("")
			assert.NoError(t, err)

			err = appCtx.SetFileExt("")
			assert.NoError(t, err)

			appCtx.SetTopic("")

			directory, err := appCtx.GetEntryDir()
			assert.NoError(t, err)
			file, err := appCtx.GetEntryFileName()
			assert.NoError(t, err)

			path := filepath.Join(directory, file)

			// Main function under test
			err = CreateNewFile(path)
			// Assert error handling
			if tt.expectedError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMsg)
			} else {
				assert.NoError(t, err)
				_, err := os.Stat(path)
				assert.NoError(t, err, "file should have been created successfully")
			}
		})
	}
}
