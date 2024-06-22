package paths

import (
	"fmt"
	"path/filepath"

	config "github.com/matthewchivers/journal/pkg/config/app"
	"github.com/matthewchivers/journal/pkg/templating"
)

// ConstructFullPath constructs the full path for the new file
// Returns format: basePath/mainPath/nestedPath
func ConstructFullPath(paths config.Paths, fileType config.FileType) (string, error) {
	basePath := paths.BaseDir
	mainPathPattern := paths.DirPattern
	if fileType.CustomDirPattern != "" {
		mainPathPattern = fileType.CustomDirPattern
	}

	nestedPathPattern := ""
	if fileType.SubDirPattern != "" {
		nestedPathPattern = fileType.SubDirPattern
	}

	mainPath, err := templating.ParsePattern(mainPathPattern, fileType)
	if err != nil {
		return "", fmt.Errorf("failed to construct base path: %w", err)
	}

	nestedPath, err := templating.ParsePattern(nestedPathPattern, fileType)
	if err != nil {
		return "", fmt.Errorf("failed to construct nested path: %w", err)
	}

	fileName, err := templating.ParsePattern(fileType.FileNamePattern, fileType)
	if err != nil {
		return "", fmt.Errorf("failed to construct file name: %w", err)
	}

	fullPath := filepath.Join(basePath, mainPath, nestedPath, fileName)
	return fullPath, nil
}
