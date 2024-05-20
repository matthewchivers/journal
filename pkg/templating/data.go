package templating

import (
	"time"

	"github.com/matthewchivers/journal/pkg/caltools"
	"github.com/matthewchivers/journal/pkg/config"
)

// TemplateData contains the template fields available for use in patterns
// patterns can be used for file name and directories in the config file
type TemplateData struct {
	// Year is the current year (e.g. 2021)
	Year string

	// Month is the current month (e.g. 01)
	Month string

	// Day is the current day (e.g. 02)
	Day string

	// WeekdayName is the name of the current day of the week (e.g. Monday)
	WeekdayName string

	// WeekdayNumber is the number of the current day of the week (e.g. 1)
	WeekdayNumber string

	// WeekCommencing is the date of the Monday of the week containing the current date (e.g. 2021-01-04)
	WeekCommencing string

	// WeekNumber is the week number of the current date (e.g. 1)
	WeekNumber string

	// FileTypeName is the name of the file type (e.g. notes/entry/diary/todo/meeting)
	FileTypeName string

	// FileExtension is the file extension of the file type (e.g. "md")
	FileExtension string
}

func PrepareTemplateData(fileType config.FileType, weekCommencing bool) (TemplateData, error) {
	timeNow := time.Now()
	data := TemplateData{
		Year:           timeNow.Format("2006"),
		Month:          timeNow.Format("01"),
		Day:            timeNow.Format("02"),
		WeekdayName:    timeNow.Weekday().String(),
		WeekdayNumber:  string(rune(timeNow.Weekday())),
		WeekCommencing: caltools.WeekCommencing(timeNow).Format("2006-01-02"),
		WeekNumber:     string(rune(caltools.WeekOfMonth(timeNow))),
		FileTypeName:   fileType.Name,
		FileExtension:  fileType.FileExtension,
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
