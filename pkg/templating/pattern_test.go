package templating

import (
	"testing"
	"time"

	"github.com/matthewchivers/journal/pkg/caltools"
	"github.com/matthewchivers/journal/pkg/config"
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
		entry        config.Entry
		want         string
		wantErr      bool
	}{
		{
			name:         "Valid template: year/month/day/",
			pathTemplate: "{{.Year}}/{{.Month}}/{{.Day}}/",
			entry:        config.Entry{},
			want:         currentYear + "/" + currentMonth + "/" + currentDay + "/",
			wantErr:      false,
		},
		{
			name:         "Valid template: year/month/week/entryID/",
			pathTemplate: "{{.Year}}/{{.Month}}/{{.WeekNumber}}/{{.EntryID}}/",
			entry:        config.Entry{ID: "foo-note"},
			want:         currentYear + "/" + currentMonth + "/" + string(rune(currentWeek)) + "/foo-note/",
			wantErr:      false,
		},
		{
			name:         "Valid template: year/month/wc weekCommencing/weekday/entryName/",
			pathTemplate: "{{.Year}}/{{.Month}}/wc {{.WeekCommencing}}/{{.WeekdayName}}/{{.EntryID}}/",
			entry:        config.Entry{ID: "foo-note"},
			want:         wcCurrentYear + "/" + wcCurrentMonth + "/wc " + weekCommencing.Format("2006-01-02") + "/" + currentWeekdayName + "/foo-note/",
			wantErr:      false,
		},
		{
			name:         "Valid template: year/month/week-commencing/template/",
			pathTemplate: "{{.Year}}/{{.Month}}/{{.WeekCommencing}}/{{.EntryID}}/",
			entry:        config.Entry{ID: "foo-note"},
			want:         currentYear + "/" + currentMonth + "/" + weekCommencing.Format("2006-01-02") + "/foo-note/",
			wantErr:      false,
		},
		{
			name:         "Path template without placeholders",
			pathTemplate: "static-path/foo-note.md",
			entry:        config.Entry{ID: "foo-note"},
			want:         "static-path/foo-note.md",
			wantErr:      false,
		},
		{
			name:         "Invalid path template",
			pathTemplate: "{{.Year}/{{.Month}}/{{.Day}}/{{.EntryID}}/", // Missing a closing brace
			entry:        config.Entry{ID: "foo-note"},
			want:         "",
			wantErr:      true,
		},
		{
			name:         "Empty path template",
			pathTemplate: "",
			entry:        config.Entry{},
			want:         "",
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePattern(tt.pathTemplate, tt.entry)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
