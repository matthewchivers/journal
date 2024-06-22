package templating

import (
	"time"

	"github.com/matthewchivers/journal/pkg/caltools"
)

// TemplateData contains the template fields available for use in patterns
// patterns can be used for file name and directories in the config file
type TemplateData struct {
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
}

// PrepareTemplateData creates a new TemplateData struct with the current date and file type
// If weekCommencing is true, the WeekCommencing date is used to calculate year/month/week number
func PrepareTemplateData(entryID string, fileExtension string, weekCommencing bool) (TemplateData, error) {
	timeNow := time.Now()
	data := TemplateData{
		Year:             timeNow.Format("2006"),
		YearShort:        timeNow.Format("06"),
		Month:            timeNow.Format("01"),
		MonthName:        timeNow.Month().String(),
		MonthNameShort:   timeNow.Month().String()[:3],
		Day:              timeNow.Format("02"),
		DayOrdinal:       caltools.OrdinalSuffix(timeNow.Day()),
		WeekdayName:      timeNow.Weekday().String(),
		WeekdayNameShort: timeNow.Weekday().String()[:3],
		WeekdayNumber:    string(rune(timeNow.Weekday())),
		WeekCommencing:   caltools.WeekCommencing(timeNow).Format("2006-01-02"),
		WeekNumber:       string(rune(caltools.WeekOfMonth(timeNow))),
		EntryID:          entryID,
		FileExtension:    fileExtension,
	}

	// WeekCommencing directories should nest in the same Year/Month as the commencing date.
	// e.g. 1st May 2024 is in the 5th Month, but the week commencing is 29th April 2024, so it should be in the 4th Month directory
	if weekCommencing {
		wc := caltools.WeekCommencing(timeNow)
		data.Year = wc.Format("2006")
		data.Month = wc.Format("01")
		data.WeekNumber = string(rune(caltools.WeekOfMonth(wc)))
	}

	return data, nil
}
