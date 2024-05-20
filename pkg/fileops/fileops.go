package fileops

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/matthewchivers/journal/pkg/config"
	"github.com/matthewchivers/journal/pkg/templating"
)

// CreateNewFile creates a new file based on the provided configuration and document template name
func CreateNewFile(filePath string) error {
	if err := ensureDirectoryExists(filepath.Dir(filePath)); err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	fmt.Printf("New file created: %s\n", filePath)
	return nil
}

// ensureDirectoryExists checks if the directory exists, and creates it if it does not
func ensureDirectoryExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory(s): %w", err)
		}
	}
	return nil
}

// GetFileName returns the file name for the file type
func GetFileName(fileType config.FileType) string {
	fileNameRaw := fileType.FileNamePattern
	if fileNameRaw == "" {
		fileNameRaw = fileType.Name
	}

	fileNameParsed, err := templating.ParsePattern(fileNameRaw, fileType)
	if err != nil {
		return fmt.Sprintf("failed to parse file name: %s", err)
	}

	return fileNameParsed
}
