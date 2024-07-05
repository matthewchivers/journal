package paths

import (
	"os"
	"path/filepath"
)

var (
	// AppHomePath is the path to the application
	AppHomePath string
)

// GetAppHomePath returns the path to the application
func GetAppHomePath() (string, error) {
	if AppHomePath != "" {
		return AppHomePath, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	appHomePath := filepath.Join(home, ".journal")
	return appHomePath, nil
}
