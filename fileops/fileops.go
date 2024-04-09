package fileops

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/matthewchivers/journal/config"
)

func CreateNewFile(cfg *config.Config, docTemplateName string) error {
	baseDirectory := cfg.Paths.JournalBaseDir

	nestedPath, err := ConstructPath(cfg.DocumentNestingPath, docTemplateName)
	if err != nil {
		return fmt.Errorf("failed to construct nested path: %w", err)
	}

	fileName := fmt.Sprintf("%s.md", docTemplateName)

	fullPath := filepath.Join(baseDirectory, nestedPath, fileName)

	dirPath := filepath.Dir(fullPath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory(s): %w", err)
		}
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("ConstructPath: %w", err)
	}
	defer file.Close()
	fmt.Printf("New file created: %s\n", fullPath)
	return nil
}
