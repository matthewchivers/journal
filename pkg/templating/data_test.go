package templating

import (
	"testing"
	"time"

	"github.com/matthewchivers/journal/pkg/caltools"
	"github.com/matthewchivers/journal/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestPrepareTemplateData(t *testing.T) {
	date := time.Now()
	// weekCommencing := caltools.WeekCommencing(date)
	currentWeek := caltools.WeekOfMonth(date)
	currentWeekdayName := date.Weekday().String()

	currentYear := date.Format("2006")
	currentMonth := date.Format("01")
	currentDay := date.Format("02")

	// wcCurrentYear := weekCommencing.Format("2006")
	// wcCurrentMonth := weekCommencing.Format("01")

	type args struct {
		fileType       config.FileType
		weekCommencing bool
	}
	tests := []struct {
		name    string
		args    args
		want    TemplateData
		wantErr bool
	}{
		{
			name: "Test PrepareTemplateData",
			args: args{
				fileType: config.FileType{Name: "notes", FileExtension: "md"},
			},
			want: TemplateData{
				Year:           currentYear,
				Month:          currentMonth,
				Day:            currentDay,
				WeekdayName:    currentWeekdayName,
				WeekdayNumber:  string(rune(date.Weekday())),
				WeekCommencing: caltools.WeekCommencing(date).Format("2006-01-02"),
				WeekNumber:     string(rune(currentWeek)),
				FileTypeName:   "notes",
				FileExtension:  "md",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := PrepareTemplateData(tt.args.fileType, tt.args.weekCommencing)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, actual)
			}
		})
	}
}
