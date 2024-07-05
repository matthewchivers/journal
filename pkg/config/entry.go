package config

// Entry contains the configuration for a entry type
type Entry struct {
	// ID is the identifier for the entry
	ID string `yaml:"id"`

	// FileExt is the file extension to use when creating a new entry
	FileExt string `yaml:"fileExt,omitempty"`

	// Schedule contains the schedule for the file type
	Schedule Schedule `yaml:"schedule,omitempty"`

	// JournalDirectory is the pattern to use when creating a subdirectory
	Directory string `yaml:"directory,omitempty"`

	// FileName is the pattern to use when creating a file name
	FileName string `yaml:"fileName"`

	// JournalDirOverride is an override pattern for the main directory
	JournalDirOverride string `yaml:"journalDirOverride,omitempty"`

	// CustomBaseDirectory
	BaseDirectoryOverride string `yaml:"baseDirectoryOverride,omitempty"`

	// TemplateName is the name of the template to use when creating a new entry
	// (if not specified, the default template will be used)
	TemplateName string `yaml:"templateName,omitempty"`

	// Topic is a name to be used for templating (e.g. a meeting about a certain topic)
	// Expect this to be primarily set using cli params, but can be set in the config file
	Topic string `yaml:"topic,omitempty"`
}
