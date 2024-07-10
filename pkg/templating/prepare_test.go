package templating

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPopulateDate(t *testing.T) {
	testTime := time.Date(2024, 6, 28, 0, 0, 0, 0, time.UTC)
	date := PopulateDate(testTime)

	assert.Equal(t, "2024", date.Year.Num)
	assert.Equal(t, "24", date.Year.Short)
	assert.Equal(t, "6", date.Month.Num)
	assert.Equal(t, "28", date.Day.Num)
}

func TestParsePattern(t *testing.T) {
	testTime := time.Date(2024, 6, 28, 0, 0, 0, 0, time.UTC)

	type args struct {
		entryID string
		fileExt string
		topic   string
	}

	tests := []struct {
		name     string
		args     args
		pattern  string
		expected string
	}{
		{
			name:     "base",
			args:     args{},
			pattern:  "/journal/{{.Year.Num}}/{{.Month.Num}}/{{.Day.Num}}",
			expected: "/journal/2024/6/28",
		},
		{
			name:     "week commencing",
			pattern:  "/journal/{{.WkCom.Year.Num}}/{{.WkCom.Month.Num}}/{{.WkCom.Day.Num}}",
			expected: "/journal/2024/6/24",
		},
		{
			name: "entry id / other args",
			args: args{
				entryID: "foo",
				fileExt: "md",
				topic:   "bar",
			},
			pattern:  "/journal/{{.Year.Num}}/{{.Month.Num}}/{{.Day.Num}}/{{.EntryID}}/{{.Topic}}.{{.FileExtension}}",
			expected: "/journal/2024/6/28/foo/bar.md",
		},
		{
			name:     "ordinals and other weird combinations",
			pattern:  "/journal/{{.Year.Num}}/{{.Month.Num}}/{{.Day.Num}}/{{.Month.Ord}}/{{.Year.Week.Ord}}/{{.Day.Ord}}/{{.Year.Day.Ord}}",
			expected: "/journal/2024/6/28/6th/26th/28th/180th",
		},
		{
			name:     "days in",
			pattern:  "/journal/{{.Year.Num}}/{{.Month.Ord}} of 12 months/{{.Year.Day.Ord}} day of {{.Year.DaysIn}}",
			expected: "/journal/2024/6th of 12 months/180th day of 366",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			templateData, _ := PrepareTemplateData(testTime)
			templateData.EntryID = tt.args.entryID
			templateData.FileExtension = tt.args.fileExt
			templateData.Topic = tt.args.topic
			parsedPath, err := templateData.ParsePattern(tt.pattern)
			assert.NoError(t, err, "Error should be nil")
			assert.Equal(t, tt.expected, parsedPath, "Parsed path should match expected path")
		})
	}
}

func TestPopulateWeekday(t *testing.T) {
	testTime := time.Date(2024, 6, 28, 0, 0, 0, 0, time.UTC)
	weekday := populateWeekday(testTime)
	assert.Equal(t, "5", weekday.Num)
	assert.Equal(t, "05", weekday.Pad)
	assert.Equal(t, "5th", weekday.Ord)
	assert.Equal(t, "Friday", weekday.Name)
	assert.Equal(t, "Fri", weekday.Short)
}

func TestGetYearWeek(t *testing.T) {
	testTime := time.Date(2024, 6, 28, 0, 0, 0, 0, time.UTC)
	weekDay := populateWeekday(testTime)
	yearWeek := getYearWeek(testTime, weekDay)

	assert.Equal(t, "26", yearWeek.Num)
	assert.Equal(t, "26", yearWeek.Pad)
	assert.Equal(t, "26th", yearWeek.Ord)
	assert.Equal(t, "5", yearWeek.Day.Num)
	assert.Equal(t, "05", yearWeek.Day.Pad)
	assert.Equal(t, "5th", yearWeek.Day.Ord)
	assert.Equal(t, "Friday", yearWeek.Day.Name)
	assert.Equal(t, "Fri", yearWeek.Day.Short)
}

func TestGetYearDay(t *testing.T) {
	testTime := time.Date(2024, 6, 28, 0, 0, 0, 0, time.UTC)
	yearDay := getYearDay(testTime)

	assert.Equal(t, "180", yearDay.Num)
	assert.Equal(t, "180", yearDay.Pad)
	assert.Equal(t, "180th", yearDay.Ord)
}

func TestPopulateMonth(t *testing.T) {
	testTime := time.Date(2024, 6, 28, 0, 0, 0, 0, time.UTC)
	month := PopulateMonth(testTime)

	assert.Equal(t, "6", month.Num)
	assert.Equal(t, "06", month.Pad)
	assert.Equal(t, "6th", month.Ord)
	assert.Equal(t, "June", month.Name)
	assert.Equal(t, "Jun", month.Short)
	assert.Equal(t, "30", month.DaysIn)
}
