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
	// Check if the file already exists
	if _, err := os.Stat(filePath); err == nil {
		// check if the file is a directory
		if info, err := os.Stat(filePath); err == nil && info.IsDir() {
			return fmt.Errorf("file already exists and is a directory: %s", filePath)
		}
		return fmt.Errorf("file already exists: %s", filePath)
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
