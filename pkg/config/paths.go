package config

// Paths contains the paths to directories used by the application
type Paths struct {
	// TemplatesDirectory is the path to the templates directory (default: ~/.journal/templates)
	TemplatesDirectory string `yaml:"templatesDirectory,omitempty"`

	// BaseDirectory is the base directory for entries (default: ~/journal)
	BaseDirectory string `yaml:"baseDirectory"`

	// JournalDirectory is the pattern to use when creating a directory inside the BaseDirectory directory
	JournalDirectory string `yaml:"journalDirectory,omitempty"`
}
