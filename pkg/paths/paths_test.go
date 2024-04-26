package paths

import (
	"testing"
	"time"

	"github.com/matthewchivers/journal/pkg/caltools"
	"github.com/stretchr/testify/assert"
)

func TestParsePathTemplate(t *testing.T) {
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
		name            string
		pathTemplate    string
		docTemplateName string
		want            string
		wantErr         bool
	}{
		{
			name:            "Valid template: year/month/day/",
			pathTemplate:    "{{.Year}}/{{.Month}}/{{.Day}}/",
			docTemplateName: "my-journal",
			want:            currentYear + "/" + currentMonth + "/" + currentDay + "/",
			wantErr:         false,
		},
		{
			name:            "Valid template: year/month/week/templateName/",
			pathTemplate:    "{{.Year}}/{{.Month}}/{{.WeekNumber}}/{{.TemplateName}}/",
			docTemplateName: "my-journal",
			want:            currentYear + "/" + currentMonth + "/" + string(rune(currentWeek)) + "/my-journal/",
			wantErr:         false,
		},
		{
			name:            "Valid template: year/month/wc weekCommencing/weekday/template/",
			pathTemplate:    "{{.Year}}/{{.Month}}/wc {{.WeekCommencing}}/{{.WeekdayName}}/{{.TemplateName}}/",
			docTemplateName: "my-journal",
			want:            wcCurrentYear + "/" + wcCurrentMonth + "/wc " + weekCommencing.Format("2006-01-02") + "/" + currentWeekdayName + "/my-journal/",
			wantErr:         false,
		},
		{
			name:            "Valid template: year/month/week-commencing/template/",
			pathTemplate:    "{{.Year}}/{{.Month}}/{{.WeekCommencing}}/{{.TemplateName}}/",
			docTemplateName: "my-journal",
			want:            currentYear + "/" + currentMonth + "/" + weekCommencing.Format("2006-01-02") + "/my-journal/",
			wantErr:         false,
		},
		{
			name:            "Path template without placeholders",
			pathTemplate:    "static-path/my-journal.md",
			docTemplateName: "ignored-template",
			want:            "static-path/my-journal.md",
			wantErr:         false,
		},
		{
			name:            "Invalid path template",
			pathTemplate:    "{{.Year}/{{.Month}}/{{.Day}}/{{.TemplateName}}/", // Missing a closing brace
			docTemplateName: "my-journal",
			want:            "",
			wantErr:         true,
		},
		{
			name:            "Empty path template",
			pathTemplate:    "",
			docTemplateName: "my-journal",
			want:            "",
			wantErr:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePathTemplate(tt.pathTemplate, tt.docTemplateName)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
