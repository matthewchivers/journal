package config

// FileType contains the configuration for a file type
type FileType struct {
	// Name is the name of the file type (not the name of the file being created)
	Name string `yaml:"name"`

	// FileExtension is the file extension to use when creating a new entry
	FileExtension string `yaml:"fileExtension,omitempty"`

	// Schedule contains the schedule for the file type
	Schedule Schedule `yaml:"schedule,omitempty"`

	// SubDirPattern is the pattern to use when creating a subdirectory
	SubDirPattern string `yaml:"subDirPattern,omitempty"`

	// FileNamePattern is the pattern to use when creating a file name
	FileNamePattern string `yaml:"fileNamePattern,omitempty"`

	// CustomDirPattern is an override pattern for the main directory
	CustomDirPattern string `yaml:"customDirPattern,omitempty"`

	// TemplateName is the name of the template to use when creating a new entry
	// (if not specified, the default template will be used)
	TemplateName string `yaml:"templateName,omitempty"`
}
