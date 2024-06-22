package app

// Entry contains the configuration for a entry type
type Entry struct {
	// ID is the identifier for the entry
	ID string `yaml:"id"`

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
