package config

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
	// {{.EntryID}}    - Name of the document/file type being created (e.g. note, todo, etc.)
	// {{.FileExtension}}   - File extension of the document/file type being created (e.g. md)
	// example: {{.Year}}/{{.Month}}/{{.Day}}/{{.EntryID}} -> 2024/01/02/note
	// or: {{.Year}}/{{.WeekCommencing}}/{{.WeekdayName}}/{{.EntryID}} -> 2024/2024-01-02/Monday/note
	DirPattern string `yaml:"dirPattern,omitempty"`
}
