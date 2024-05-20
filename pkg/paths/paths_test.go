package paths

import (
	"testing"
	"time"

	"github.com/matthewchivers/journal/pkg/caltools"
	"github.com/matthewchivers/journal/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestConstructFullPath(t *testing.T) {
	date := time.Now()
	weekCommencing := caltools.WeekCommencing(date)
	currentWeek := caltools.WeekOfMonth(date)
	currentWeekdayName := date.Weekday().String()

	currentYear := date.Format("2006")
	currentMonth := date.Format("01")
	currentDay := date.Format("02")

	wcCurrentYear := weekCommencing.Format("2006")
	wcCurrentMonth := weekCommencing.Format("01")

	type args struct {
		paths    config.Paths
		fileType config.FileType
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Valid template: year/month/week/FileTypeName",
			args: args{
				paths: config.Paths{
					BaseDir:    "base",
					DirPattern: "{{.Year}}/{{.Month}}/{{.WeekNumber}}/{{.FileTypeName}}",
				},
				fileType: config.FileType{Name: "foo-note"},
			},
			want:    "base/" + currentYear + "/" + currentMonth + "/" + string(rune(currentWeek)) + "/foo-note",
			wantErr: false,
		},
		{
			name: "Valid template: year/month/wc weekCommencing/weekday/template",
			args: args{
				paths: config.Paths{
					BaseDir:    "base",
					DirPattern: "{{.Year}}/{{.Month}}/wc {{.WeekCommencing}}/{{.WeekdayName}}/{{.FileTypeName}}",
				},
				fileType: config.FileType{Name: "foo-note"},
			},
			want:    "base/" + wcCurrentYear + "/" + wcCurrentMonth + "/wc " + weekCommencing.Format("2006-01-02") + "/" + currentWeekdayName + "/foo-note",
			wantErr: false,
		},
		{
			name: "Valid template: year/month/week-commencing/template",
			args: args{
				paths: config.Paths{
					BaseDir:    "base",
					DirPattern: "{{.Year}}/{{.Month}}/{{.WeekCommencing}}/{{.FileTypeName}}",
				},
				fileType: config.FileType{Name: "foo-note"},
			},
			want:    "base/" + currentYear + "/" + currentMonth + "/" + weekCommencing.Format("2006-01-02") + "/foo-note",
			wantErr: false,
		},
		{
			name: "Path template without placeholders",
			args: args{
				paths: config.Paths{
					BaseDir:    "base",
					DirPattern: "static-path/foo-note.md",
				},
				fileType: config.FileType{Name: "foo-note"},
			},
			want:    "base/static-path/foo-note.md",
			wantErr: false,
		},
		{
			name: "Path template with custom file path",
			args: args{
				paths: config.Paths{
					BaseDir:    "base",
					DirPattern: "static-path/{{.FileTypeName}}",
				},
				fileType: config.FileType{
					Name:             "foo-note",
					CustomDirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
					FileNamePattern:  "{{.FileTypeName}}.{{.FileExtension}}",
					FileExtension:    "md",
				},
			},
			want:    "base/" + currentYear + "/" + currentMonth + "/" + currentDay + "/foo-note.md",
			wantErr: false,
		},
		{
			name: "Custom subdir pattern",
			args: args{
				paths: config.Paths{
					BaseDir:    "base",
					DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
				},
				fileType: config.FileType{
					Name:          "foo-note",
					SubDirPattern: "{{.FileTypeName}}s",
				},
			},
			want: "base/" + currentYear + "/" + currentMonth + "/" + currentDay + "/" + "foo-notes",
		},
		{
			name: "Custom dir and subdir pattern",
			args: args{
				paths: config.Paths{
					BaseDir:    "base",
					DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
				},
				fileType: config.FileType{
					Name:             "foo-note",
					CustomDirPattern: "{{.Year}}/{{.Month}}/{{.WeekCommencing}}",
					SubDirPattern:    "{{.FileTypeName}}s",
				},
			},
			want: "base/" + currentYear + "/" + currentMonth + "/" + weekCommencing.Format("2006-01-02") + "/" + "foo-notes",
		},
		{
			name: "Custom file name pattern",
			args: args{
				paths: config.Paths{
					BaseDir:    "base",
					DirPattern: "{{.Year}}/{{.Month}}/{{.WeekCommencing}}",
				},
				fileType: config.FileType{
					Name:            "foo-note",
					SubDirPattern:   "{{.FileTypeName}}s",
					FileExtension:   "md",
					FileNamePattern: "{{.FileTypeName}}-{{.Day}}-{{.Month}}-{{.Year}}.{{.FileExtension}}",
				},
			},
			want: "base/" + currentYear + "/" + currentMonth + "/" + weekCommencing.Format("2006-01-02") + "/" + "foo-notes/foo-note-" + currentDay + "-" + currentMonth + "-" + currentYear + ".md",
		},
		{
			name: "Invalid path template",
			args: args{
				paths: config.Paths{
					BaseDir:    "base",
					DirPattern: "{{.Year}/{{.Month}}/{{.Day}}/{{.FileTypeName}}",
				},
				fileType: config.FileType{Name: "foo-note"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ConstructFullPath(tt.args.paths, tt.args.fileType)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, actual)
			}
		})
	}
}
