package caltools

import (
	"time"
)

// WeekOfMonth calculates the week of the month for the given date.
func WeekOfMonth(t time.Time) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())

	offset := (int(firstOfMonth.Weekday()) + 6) % 7
	week := (t.Day() + offset) / 7

	// Account for partial weeks at beginning of month
	// e.g. If the month started on a Wednesday, the calculation (without adjustment)
	// would still put you in the first week, because (6 days + 2 offset / 7 = 1.14,
	// which gets rounded down to 1). But realistically the 6th is in the second week
	// of the month.
	if offset == 0 || (t.Day()+offset)%7 != 0 {
		week++
	}

	return week
}

// WeekCommencing calculates the date of the Monday of the week containing the given date.
func WeekCommencing(t time.Time) time.Time {
	offset := (int(t.Weekday()) + 6) % 7
	weekCommencing := t.AddDate(0, 0, offset*-1)
	return weekCommencing
}

// OrdinalSuffix returns the ordinal suffix for the given day of the month.
// e.g. 1st, 2nd, 3rd, 4th
func OrdinalSuffix(day int) string {
	if day >= 11 && day <= 13 {
		return "th"
	}

	switch day % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

func DaysInMonth(t time.Time) int {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, t.Location()).Day()
}

// DaysInYear returns the number of days in the given year.
func DaysInYear(t time.Time) int {
	return time.Date(t.Year()+1, 1, 0, 0, 0, 0, 0, t.Location()).YearDay()
}
