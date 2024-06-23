package templating

import (
	"strings"
	"testing"
	"time"

	"github.com/matthewchivers/journal/pkg/caltools"
	"github.com/stretchr/testify/assert"
)

func TestPrepareTemplateData(t *testing.T) {
	date := time.Now()
	// weekCommencing := caltools.WeekCommencing(date)
	currentWeek := caltools.WeekOfMonth(date)
	currentWeekdayName := date.Weekday().String()

	currentYear := date.Format("2006")
	currentYearShort := date.Format("06")
	currentMonth := date.Format("01")
	currentDay := date.Format("02")

	// wcCurrentYear := weekCommencing.Format("2006")
	// wcCurrentMonth := weekCommencing.Format("01")

	tests := []struct {
		name    string
		want    TemplateModel
		wantErr bool
	}{
		{
			name: "Test PrepareTemplateData",
			want: TemplateModel{
				Year:             currentYear,
				YearShort:        currentYearShort,
				Month:            currentMonth,
				MonthName:        date.Month().String(),
				MonthNameShort:   date.Month().String()[:3],
				Day:              currentDay,
				DayOrdinal:       caltools.OrdinalSuffix(date.Day()),
				WeekdayName:      currentWeekdayName,
				WeekdayNameShort: currentWeekdayName[:3],
				WeekdayNumber:    string(rune(date.Weekday())),
				WeekCommencing:   caltools.WeekCommencing(date).Format("2006-01-02"),
				WeekNumber:       string(rune(currentWeek)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := PrepareTemplateData(date)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, actual)
			}
		})
	}
}

func TestParsePattern(t *testing.T) {
	date := time.Now()
	weekCommencing := caltools.WeekCommencing(date)
	currentWeek := caltools.WeekOfMonth(date)
	currentWeekdayName := date.Weekday().String()

	currentYear := date.Format("2006")
	currentMonth := date.Format("01")
	currentDay := date.Format("02")

	wcCurrentYear := weekCommencing.Format("2006")
	wcCurrentMonth := weekCommencing.Format("01")

	tests := []struct {
		name          string
		pattern       string
		entryID       string
		fileExtension string
		want          string
		wantErr       bool
	}{
		{
			name:    "Valid template: year/month/day/",
			pattern: "{{.Year}}/{{.Month}}/{{.Day}}/",
			entryID: "",
			want:    currentYear + "/" + currentMonth + "/" + currentDay + "/",
			wantErr: false,
		},
		{
			name:    "Valid template: year/month/week/entryID/",
			pattern: "{{.Year}}/{{.Month}}/{{.WeekNumber}}/{{.EntryID}}/",
			entryID: "foo-note",
			want:    currentYear + "/" + currentMonth + "/" + string(rune(currentWeek)) + "/foo-note/",
			wantErr: false,
		},
		{
			name:    "Valid template: year/month/wc weekCommencing/weekday/entryName/",
			pattern: "{{.Year}}/{{.Month}}/wc {{.WeekCommencing}}/{{.WeekdayName}}/{{.EntryID}}/",
			entryID: "foo-note",
			want:    wcCurrentYear + "/" + wcCurrentMonth + "/wc " + weekCommencing.Format("2006-01-02") + "/" + currentWeekdayName + "/foo-note/",
			wantErr: false,
		},
		{
			name:    "Valid template: year/month/week-commencing/template/",
			pattern: "{{.Year}}/{{.Month}}/{{.WeekCommencing}}/{{.EntryID}}/",
			entryID: "foo-note",
			want:    currentYear + "/" + currentMonth + "/" + weekCommencing.Format("2006-01-02") + "/foo-note/",
			wantErr: false,
		},
		{
			name:    "Path template without placeholders",
			pattern: "static-path/foo-note.md",
			entryID: "foo-note",
			want:    "static-path/foo-note.md",
			wantErr: false,
		},
		{
			name:    "Invalid path template",
			pattern: "{{.Year}/{{.Month}}/{{.Day}}/{{.EntryID}}/", // Missing a closing brace
			entryID: "foo-note",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Empty path template",
			pattern: "",
			entryID: "",
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			templateModel, err := PrepareTemplateData(date)
			assert.NoError(t, err)

			templateModel.EntryID = tt.entryID
			templateModel.FileExtension = "md"

			if strings.Contains(tt.pattern, "{{.WeekCommencing}}") {
				templateModel.AdjustForWeekCommencing(date)
			}

			got, err := templateModel.ParsePattern(tt.pattern)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
