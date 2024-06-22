package app

// Config contains the configuration for the application
type Config struct {
	// DefaultFileType is the default file type (name) to use when creating a new entry
	DefaultFileType string `yaml:"defaultFileType"`

	// DefaultFileExtension is the default file extension to use when creating a new entry
	DefaultFileExtension string `yaml:"defaultFileExtension,omitempty"`

	// FileTypes is a list of file types
	FileTypes []FileType `yaml:"fileTypes"`

	// Paths contains the paths to directories used by the application
	Paths Paths `yaml:"paths"`

	// UserSettings contains user-specific settings
	UserSettings UserSettings `yaml:"userSettings,omitempty"`
}
