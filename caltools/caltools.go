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
