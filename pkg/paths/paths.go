package paths

import (
	"os"
	"path/filepath"
)

var (
	appHomePath string
)

// GetAppHomePath returns the path to the application
func GetAppHomePath() (string, error) {
	if appHomePath != "" {
		return appHomePath, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	appHomePath := filepath.Join(home, ".journal")
	return appHomePath, nil
}

// SetAppHomePath sets the application home path
func SetAppHomePath(path string) {
	appHomePath = path
}
