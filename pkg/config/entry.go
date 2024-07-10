package config

// Entry contains the configuration for a entry type
type Entry struct {
	// ID is the identifier for the entry
	ID string `yaml:"id"`

	// FileExtension is the file extension to use when creating a new entry
	FileExtension string `yaml:"fileExt,omitempty"`

	// Schedule contains the schedule for the file type
	Schedule Schedule `yaml:"schedule,omitempty"`

	// DirectoryPattern is the pattern to use when creating a subdirectory
	DirectoryPattern string `yaml:"directoryPattern,omitempty"`

	// FileNamePattern is the pattern to use when creating a file name
	FileNamePattern string `yaml:"fileNamePattern"`

	// BaseDirectory
	BaseDirectory string `yaml:"baseDirectory,omitempty"`

	// TemplateName is the name of the template to use when creating a new entry
	// (if not specified, the default template will be used)
	TemplateName string `yaml:"templateName,omitempty"`

	// Topic is a name to be used for templating (e.g. a meeting about a certain topic)
	// Expect this to be primarily set using cli params, but can be set in the config file
	Topic string `yaml:"topic,omitempty"`

	// Editor is the editor to use when opening files
	Editor string `yaml:"editor,omitempty"`
}
