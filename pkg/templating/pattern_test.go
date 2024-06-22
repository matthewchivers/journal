package templating

import (
	"testing"
	"time"

	"github.com/matthewchivers/journal/pkg/caltools"
	"github.com/matthewchivers/journal/pkg/config/app"
	"github.com/stretchr/testify/assert"
)

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
		name         string
		pathTemplate string
		fileType     app.FileType
		want         string
		wantErr      bool
	}{
		{
			name:         "Valid template: year/month/day/",
			pathTemplate: "{{.Year}}/{{.Month}}/{{.Day}}/",
			fileType:     app.FileType{},
			want:         currentYear + "/" + currentMonth + "/" + currentDay + "/",
			wantErr:      false,
		},
		{
			name:         "Valid template: year/month/week/FileTypeName/",
			pathTemplate: "{{.Year}}/{{.Month}}/{{.WeekNumber}}/{{.FileTypeName}}/",
			fileType:     app.FileType{Name: "foo-note"},
			want:         currentYear + "/" + currentMonth + "/" + string(rune(currentWeek)) + "/foo-note/",
			wantErr:      false,
		},
		{
			name:         "Valid template: year/month/wc weekCommencing/weekday/template/",
			pathTemplate: "{{.Year}}/{{.Month}}/wc {{.WeekCommencing}}/{{.WeekdayName}}/{{.FileTypeName}}/",
			fileType:     app.FileType{Name: "foo-note"},
			want:         wcCurrentYear + "/" + wcCurrentMonth + "/wc " + weekCommencing.Format("2006-01-02") + "/" + currentWeekdayName + "/foo-note/",
			wantErr:      false,
		},
		{
			name:         "Valid template: year/month/week-commencing/template/",
			pathTemplate: "{{.Year}}/{{.Month}}/{{.WeekCommencing}}/{{.FileTypeName}}/",
			fileType:     app.FileType{Name: "foo-note"},
			want:         currentYear + "/" + currentMonth + "/" + weekCommencing.Format("2006-01-02") + "/foo-note/",
			wantErr:      false,
		},
		{
			name:         "Path template without placeholders",
			pathTemplate: "static-path/foo-note.md",
			fileType:     app.FileType{Name: "foo-note"},
			want:         "static-path/foo-note.md",
			wantErr:      false,
		},
		{
			name:         "Invalid path template",
			pathTemplate: "{{.Year}/{{.Month}}/{{.Day}}/{{.FileTypeName}}/", // Missing a closing brace
			fileType:     app.FileType{Name: "foo-note"},
			want:         "",
			wantErr:      true,
		},
		{
			name:         "Empty path template",
			pathTemplate: "",
			fileType:     app.FileType{},
			want:         "",
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePattern(tt.pathTemplate, tt.fileType)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
