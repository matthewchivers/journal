package fileops

import (
	"os"
	"testing"

	"github.com/matthewchivers/journal/pkg/config"
	"github.com/matthewchivers/journal/pkg/paths"
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
				FileTypes: []config.FileType{
					{
						Name:            "foo",
						FileExtension:   "md",
						FileNamePattern: "{{.FileTypeName}}.{{.FileExtension}}",
					},
				},
			},
			expectedError: false,
		},
		{
			name: "successful file creation - hardcoded extension",
			cfg: &config.Config{
				Paths: config.Paths{
					BaseDir:    tempdir,
					DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
				},
				FileTypes: []config.FileType{
					{
						Name:            "foo",
						FileNamePattern: "{{.FileTypeName}}.md",
					},
				},
			},
			expectedError: false,
		},
		{
			name: "successful file creation - no extension",
			cfg: &config.Config{
				Paths: config.Paths{
					BaseDir:    tempdir,
					DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
				},
				FileTypes: []config.FileType{
					{
						Name:            "foo",
						FileNamePattern: "{{.FileTypeName}}",
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
				FileTypes: []config.FileType{
					{
						Name:            "foo",
						FileNamePattern: "{{.FileTypeName}}.{{.FileExtension}}",
						SubDirPattern:   "{{.FileTypeName}}s",
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
			fileType := tt.cfg.FileTypes[0]

			fullPath, err := paths.ConstructFullPath(tt.cfg.Paths, fileType)
			if err != nil {
				panic(err)
			}

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
