package paths

import (
	"fmt"
	"path/filepath"

	"github.com/matthewchivers/journal/pkg/config"
	"github.com/matthewchivers/journal/pkg/templating"
)

// ConstructFullPath constructs the full path for the new file
// Returns format: basePath/mainPath/nestedPath
func ConstructFullPath(paths config.Paths, entry config.Entry) (string, error) {
	basePath := paths.BaseDir
	mainPathPattern := paths.DirPattern
	if entry.CustomDirPattern != "" {
		mainPathPattern = entry.CustomDirPattern
	}

	nestedPathPattern := ""
	if entry.SubDirPattern != "" {
		nestedPathPattern = entry.SubDirPattern
	}

	mainPath, err := templating.ParsePattern(mainPathPattern, entry)
	if err != nil {
		return "", fmt.Errorf("failed to construct base path: %w", err)
	}

	nestedPath, err := templating.ParsePattern(nestedPathPattern, entry)
	if err != nil {
		return "", fmt.Errorf("failed to construct nested path: %w", err)
	}

	fileName, err := templating.ParsePattern(entry.FileNamePattern, entry)
	if err != nil {
		return "", fmt.Errorf("failed to construct file name: %w", err)
	}

	fullPath := filepath.Join(basePath, mainPath, nestedPath, fileName)
	return fullPath, nil
}
