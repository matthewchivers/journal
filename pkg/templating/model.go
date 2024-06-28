package templating

type Day struct {
	// Number of the day
	Num string

	// Pad is the current day zero padded (e.g. 02)
	Pad string

	// Ord is the current day with ordinal suffix (e.g. 1st, 2nd, 3rd, 4th, ... 30th, 31st)
	Ord string
}

type WeekDay struct {
	// Number of the day
	Num string

	// Pad is the current day zero padded (e.g. 02)
	Pad string

	// Ord is the current day with ordinal suffix (e.g. 1st, 2nd, 3rd, 4th, ... 30th, 31st)
	Ord string

	// Name is the name of the current day (e.g. Monday)
	Name string

	// Short is the short name of the current day (e.g. Mon)
	Short string
}

type Week struct {
	// Number of the week
	Num string

	// Pad is the current week zero padded (e.g. 01)
	Pad string

	// Ord is the current week with ordinal suffix (e.g. 1st, 2nd, 3rd, 4th, ... 52nd, 53rd)
	Ord string

	// Day is the current day of the week
	Day WeekDay
}

type Month struct {
	// Number of the month
	Num string

	// Pad is the current month zero padded (e.g. 01)
	Pad string

	// Ord is the current month with ordinal suffix (e.g. 1st, 2nd, 3rd, 4th, ... 11th, 12th)
	Ord string

	// Name is the name of the current month (e.g. January)
	Name string

	// Short is the short name of the current month (e.g. Jan)
	Short string

	// DaysIn is the number of days in the month
	DaysIn string

	// Day is the current day of the month (e.g. 1, 2, ..., 30, 31)
	Day Day

	// Week is the current week of the month
	Week Week
}

type Year struct {
	// Number of the year
	Num string

	// Short is the current year in short form (e.g. 21)
	Short string

	// Month is the current month
	Month Month

	// Week is the current week of the year
	Week Week

	// Day is the current day of the year
	Day Day

	// DaysIn is the number of days in the year
	DaysIn string
}

type Date struct {
	// Year is the current year
	Year Year

	// Month is the current month
	Month Month

	// Day is the current day of the month
	Day WeekDay
}

// TemplateModel contains the template fields available for use in patterns
// patterns can be used for file name and directories in the config file
type TemplateModel struct {
	// Date contains the current date
	Date Date

	// Week Commencing date contains the date of the Monday of the week containing the current date
	WCDate Date

	// EntryID is the name of the entry type (e.g. notes/entry/diary/todo/meeting)
	EntryID string

	// FileExt is the file extension of the file type (e.g. "md")
	FileExt string

	// Topic is the name of the topic for the entry (e.g. "project A/B/C")
	Topic string
}
