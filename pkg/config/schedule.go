package config

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
