package templating

import (
	"bytes"
	"html/template"
	"time"

	"github.com/matthewchivers/journal/pkg/caltools"
)

// TemplateModel contains the template fields available for use in patterns
// patterns can be used for file name and directories in the config file
type TemplateModel struct {
	// Year is the current year (e.g. 2021)
	Year string

	// YearShort is the current year in short form (e.g. 21)
	YearShort string

	// Month is the current month (e.g. 01)
	Month string

	// MonthName is the name of the current month (e.g. January)
	MonthName string

	// MonthNameShort is the short name of the current month (e.g. Jan)
	MonthNameShort string

	// Day is the current day (e.g. 02)
	Day string

	// Day with ordinal suffix (e.g. 1st, 2nd, 3rd, 4th)
	DayOrdinal string

	// WeekdayName is the name of the current day of the week (e.g. Monday)
	WeekdayName string

	// WeekdayNameShort is the short name of the current day of the week (e.g. Mon)
	WeekdayNameShort string

	// WeekdayNumber is the number of the current day of the week (e.g. 1)
	WeekdayNumber string

	// WeekCommencing is the date of the Monday of the week containing the current date (e.g. 2021-01-04)
	WeekCommencing string

	// WeekNumber is the week number of the current date (e.g. 1)
	WeekNumber string

	// EntryID is the name of the entry type (e.g. notes/entry/diary/todo/meeting)
	EntryID string

	// FileExtension is the file extension of the file type (e.g. "md")
	FileExtension string

	// Topic is the name of the topic for the entry (e.g. "project A/B/C")
	Topic string
}

// PrepareTemplateData creates a new TemplateModel struct with the current date and file type
// If weekCommencing is true, the WeekCommencing date is used to calculate year/month/week number
func PrepareTemplateData(time time.Time) (TemplateModel, error) {
	data := TemplateModel{
		Year:             time.Format("2006"),
		YearShort:        time.Format("06"),
		Month:            time.Format("01"),
		MonthName:        time.Month().String(),
		MonthNameShort:   time.Month().String()[:3],
		Day:              time.Format("02"),
		DayOrdinal:       caltools.OrdinalSuffix(time.Day()),
		WeekdayName:      time.Weekday().String(),
		WeekdayNameShort: time.Weekday().String()[:3],
		WeekdayNumber:    string(rune(time.Weekday())),
		WeekCommencing:   caltools.WeekCommencing(time).Format("2006-01-02"),
		WeekNumber:       string(rune(caltools.WeekOfMonth(time))),
	}
	return data, nil
}

func (tm *TemplateModel) AdjustForWeekCommencing(time time.Time) {
	time = caltools.WeekCommencing(time)
	time = caltools.WeekCommencing(time)
	tm.Year = time.Format("2006")
	tm.Month = time.Format("01")
	tm.WeekNumber = string(rune(caltools.WeekOfMonth(time)))
}

// ParsePattern creates a new path for a journal entry based on a path template
func (tm *TemplateModel) ParsePattern(pattern string) (string, error) {
	t, err := template.New("path").Parse(pattern)
	if err != nil {
		return "", err
	}

	var templateB bytes.Buffer
	err = t.Execute(&templateB, tm)
	if err != nil {
		return "", err
	}

	parsedTemplate := templateB.String()

	return parsedTemplate, nil
}
