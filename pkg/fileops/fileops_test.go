package fileops

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	app "github.com/matthewchivers/journal/pkg/application"
	"github.com/matthewchivers/journal/pkg/config"
	"github.com/matthewchivers/journal/pkg/logger"
	"github.com/rs/zerolog"
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
		expectedFilePath string
		expectedError    bool
		expectedErrorMsg string
	}{
		{
			name: "successful file creation",
			cfg: &config.Config{
				Paths: config.Paths{
					BaseDirectory:    tempdir,
					JournalDirectory: "{{.Year.Num}}/{{.Month.Pad}}/{{.Day.Pad}}",
				},
				Entries: []config.Entry{
					{
						ID:       "foo",
						FileExt:  "md",
						FileName: "{{.EntryID}}.{{.FileExt}}",
					},
				},
			},
			expectedFilePath: filepath.Join(tempdir, time.Now().Format("2006/01/02")+"/foo.md"),
			expectedError:    false,
		},
		{
			name: "successful file creation - hardcoded extension", // support this as valid functionality in the future
			cfg: &config.Config{
				Paths: config.Paths{
					BaseDirectory:    tempdir,
					JournalDirectory: "{{.Year.Num}}/{{.Month.Pad}}/{{.Day.Pad}}",
				},
				Entries: []config.Entry{
					{
						ID:       "foo",
						FileName: "{{.EntryID}}.md",
					},
				},
			},
			expectedFilePath: filepath.Join(tempdir, time.Now().Format("2006/01/02")+"/foo.md"),
			expectedError:    true,
		},
		{
			name: "successful file creation - no extension",
			cfg: &config.Config{
				Paths: config.Paths{
					BaseDirectory:    tempdir,
					JournalDirectory: "{{.Year.Num}}/{{.Month.Pad}}/{{.Day.Pad}}",
				},
				Entries: []config.Entry{
					{
						ID:       "foo",
						FileName: "{{.EntryID}}",
					},
				},
			},
			expectedFilePath: filepath.Join(tempdir, time.Now().Format("2006/01/02")+"/foo"),
			expectedError:    false,
		},
		{
			name: "successful file creation - custom subdirectory",
			cfg: &config.Config{
				Paths: config.Paths{
					BaseDirectory:    tempdir,
					JournalDirectory: "{{.Year.Num}}/{{.Month.Pad}}/{{.Day.Pad}}",
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
			expectedFilePath: filepath.Join(tempdir, time.Now().Format("2006/01/02")+"/foos/foo.md"),
			expectedError:    false,
		},
	}

	defer func() {
		// Clean up the file system
		os.RemoveAll(tempdir)
	}()

	for _, tt := range tests {
		tempLogger := zerolog.New(os.Stdout).With().Timestamp().Logger()
		logger.SetLogger(&tempLogger)
		t.Run(tt.name, func(t *testing.T) {
			entry := tt.cfg.Entries[0]

			appCtx, err := app.NewApp()
			assert.NoError(t, err)
			tt.cfg.FileExt = "md"
			appCtx.Config = tt.cfg

			appCtx.SetLaunchTime(time.Now())

			err = appCtx.PreparePatternData()
			assert.NoError(t, err)

			err = appCtx.SetEntryID(entry.ID)
			assert.NoError(t, err)

			err = appCtx.SetFileExt("")
			assert.NoError(t, err)

			err = appCtx.SetFileName("")
			assert.NoError(t, err)

			err = appCtx.SetDirectory("")
			assert.NoError(t, err)

			appCtx.SetTopic("")

			path, err := appCtx.GetFilePath()
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedFilePath, path)

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
