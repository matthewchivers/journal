package config

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

// Schedule contains the schedule for a file type
type Schedule struct {
	// Frequency is the frequency of the schedule (daily, weekly, monthly, yearly)
	Frequency string `yaml:"frequency"`

	// Interval is the interval for the schedule (e.g. Interval: 2, Frequency: monthly => every 2 months)
	Interval int `yaml:"interval,omitempty"`

	// Days is the days of the week to create entries (e.g. 1, 3, 5 => Monday, Wednesday, Friday)
	Days []int `yaml:"days,omitempty"`

	// Dates is the dates of the month to create entries (e.g. 1, 15, 30 => 1st, 15th, 30th)
	Dates []int `yaml:"dates,omitempty"`

	// Weeks are the weeks of the month to create entries (e.g. 1, 3 => 1st and 3rd weeks)
	Weeks []int `yaml:"weeks,omitempty"`

	// Months are the months of the year to create entries (e.g. 1, 3, 5 => January, March, May)
	Months []int `yaml:"months,omitempty"`
}

// Paths contains the paths to directories used by the application
type Paths struct {
	// TemplatesDir is the path to the templates directory (default: ~/.journal/templates)
	TemplatesDir string `yaml:"templatesDir,omitempty"`

	// BaseDir is the base directory for entries (default: ~/journal)
	BaseDir string `yaml:"baseDir"`

	// DirPattern is the pattern to use when creating a directory inside the BaseDir directory
	// {{.Year}}           - Current year (YYYY -> 2024)
	// {{.YearShort}}.     - Current year in short form (YY -> 24)
	// {{.Month}}          - Current month (MM -> 01)
	// {{.MonthName}}      - Current month name (January)
	// {{.MonthNameShort}} - Current month name in short form (Jan)
	// {{.Day}}            - Current day (DD -> 02)
	// {{.DayOrdinal}}     - Current day with ordinal suffix (1st, 2nd, 3rd, 4th)
	// {{.WeekdayName}}    - Current weekday (Monday, Tuesday...)
	// {{.WeekdayNameShort}} - Current weekday in short form (Mon, Tue...)
	// {{.WeekdayNumber}}  - Current weekday number (0-7)
	// {{.WeekCommencing}} - Date of the week commencing (e.g. Monday of the current week)
	// {{.WeekNumber}}     - Week number of the year
	// {{.FileTypeName}}    - Name of the document/file type being created (e.g. note, todo, etc.)
	// {{.FileExtension}}   - File extension of the document/file type being created (e.g. md)
	// example: {{.Year}}/{{.Month}}/{{.Day}}/{{.FileTypeName}} -> 2024/01/02/note
	// or: {{.Year}}/{{.WeekCommencing}}/{{.WeekdayName}}/{{.FileTypeName}} -> 2024/2024-01-02/Monday/note
	DirPattern string `yaml:"dirPattern,omitempty"`
}

// UserSettings contains user-specific settings
type UserSettings struct {
	// Timezone is the timezone to use for the application
	Timezone string `yaml:"timezone,omitempty"`
}
