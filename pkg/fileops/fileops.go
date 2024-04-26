package fileops

import (
	"fmt"
	"os"
	"path/filepath"
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
