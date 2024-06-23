package templating

import (
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

	type args struct {
		EntryID        string
		FileExtension  string
		weekCommencing bool
	}
	tests := []struct {
		name    string
		args    args
		want    TemplateModel
		wantErr bool
	}{
		{
			name: "Test PrepareTemplateData",
			args: args{EntryID: "notes", FileExtension: "md", weekCommencing: false},
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
				EntryID:          "notes",
				FileExtension:    "md",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := PrepareTemplateData(tt.args.EntryID, tt.args.FileExtension, tt.args.weekCommencing)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, actual)
			}
		})
	}
}
