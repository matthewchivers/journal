package templating

import (
	"bytes"
	"fmt"
	"html/template"
	"time"

	"github.com/matthewchivers/journal/pkg/caltools"
)

// PrepareTemplateData creates a new TemplateModel struct with the current date and file type
func PrepareTemplateData(time time.Time) (TemplateModel, error) {
	weekCommencing := caltools.WeekCommencing(time)

	data := TemplateModel{
		Year:  PopulateYear(time),
		Month: PopulateMonth(time),
		Day:   PopulateDay(time),
		WkCom: PopulateDate(weekCommencing),
	}
	return data, nil
}

// PopulateDate creates a new Date struct with the current date
func PopulateDate(time time.Time) Date {
	date := Date{
		Year:  PopulateYear(time),
		Month: PopulateMonth(time),
		Day:   PopulateDay(time),
	}

	return date
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

// PopulateDay returns a WeekDay struct using the day of the month
func PopulateDay(time time.Time) WeekDay {
	wd := &WeekDay{
		Num:   fmt.Sprintf("%d", time.Day()),
		Pad:   fmt.Sprintf("%02d", time.Day()),
		Ord:   fmt.Sprintf("%d%s", time.Day(), caltools.OrdinalSuffix(time.Day())),
		Name:  time.Weekday().String(),
		Short: time.Weekday().String()[:3],
	}
	return *wd
}

// populateWeekday returns a WeekDay struct using the day of the week rather than the day of the month
func populateWeekday(time time.Time) WeekDay {
	weekday := &WeekDay{
		Num:   fmt.Sprintf("%d", time.Weekday()),
		Pad:   fmt.Sprintf("%02d", time.Weekday()),
		Ord:   fmt.Sprintf("%d%s", time.Weekday(), caltools.OrdinalSuffix(int(time.Weekday()))),
		Name:  time.Weekday().String(),
		Short: time.Weekday().String()[:3],
	}
	return *weekday
}

// getYearWeek populates and returns a Week struct for the year
func getYearWeek(time time.Time, weekDay WeekDay) Week {
	_, yearWeekNum := time.ISOWeek()
	week := &Week{
		Num: fmt.Sprintf("%d", yearWeekNum),
		Pad: fmt.Sprintf("%02d", yearWeekNum),
		Ord: fmt.Sprintf("%d%s", yearWeekNum, caltools.OrdinalSuffix(yearWeekNum)),
		Day: weekDay,
	}
	return *week
}

// getYearDay populates and returns a Day struct for the year
func getYearDay(time time.Time) Day {
	yearDayNum := time.YearDay()
	yearDay := &Day{
		Num: fmt.Sprintf("%d", yearDayNum),
		Pad: fmt.Sprintf("%03d", yearDayNum),
		Ord: fmt.Sprintf("%d%s", yearDayNum, caltools.OrdinalSuffix(yearDayNum)),
	}
	return *yearDay
}

// PopulateMonth populates and returns a Month struct
func PopulateMonth(time time.Time) Month {
	month := &Month{
		Num:    time.Format("1"),
		Pad:    time.Format("01"),
		Ord:    fmt.Sprintf("%s%s", time.Format("1"), caltools.OrdinalSuffix(int(time.Month()))),
		Name:   time.Month().String(),
		Short:  time.Month().String()[:3],
		DaysIn: fmt.Sprintf("%d", caltools.DaysInMonth(time)),
		Day: Day{
			Num: time.Format("2"),
			Pad: time.Format("02"),
			Ord: fmt.Sprintf("%s%s", time.Format("2"), caltools.OrdinalSuffix(time.Day())),
		},
		Week: Week{
			Num: fmt.Sprintf("%d", caltools.WeekOfMonth(time)),
			Pad: fmt.Sprintf("%02d", caltools.WeekOfMonth(time)),
			Ord: fmt.Sprintf("%d%s", caltools.WeekOfMonth(time), caltools.OrdinalSuffix(caltools.WeekOfMonth(time))),
			Day: populateWeekday(time),
		},
	}
	return *month
}

// PopulateYear populates and returns a Year struct
func PopulateYear(time time.Time) Year {
	// Get structs that will be filled
	weekDay := populateWeekday(time)
	yearDay := getYearDay(time)
	yearWeek := getYearWeek(time, weekDay)

	year := &Year{
		Num:    time.Format("2006"),
		Short:  time.Format("06"),
		Month:  PopulateMonth(time),
		Week:   yearWeek,
		Day:    yearDay,
		DaysIn: fmt.Sprintf("%d", caltools.DaysInYear(time)),
	}

	return *year
}
