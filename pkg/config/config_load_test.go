package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name     string
		yamlData string
		want     *Config
		wantErr  bool
	}{
		{
			name: "Basic Configuration",
			yamlData: `
defaultEntry: "report"
entries:
  - id: "report"
paths:
  baseDirectory: "/journals"
`,
			want: &Config{
				DefaultEntry: "report",
				Entries: []Entry{
					{
						ID: "report",
					},
				},
				Paths: Paths{
					BaseDirectory: "/journals",
				},
			},
			wantErr: false,
		},
		{
			name: "Full Configuration With Schedules",
			yamlData: `
defaultEntry: "log"
entries:
 - id: "log"
   schedule:
     frequency: "daily"
     days: [1,3,5]
   templateName: "log.tmpl"
paths:
  templatesDir: "~/.journal/customtemplates"
  baseDirectory: "~/journals"
  journalDirectory: "{{.Year}}/{{.Month}}/{{.Day}}/"
userSettings:
  timezone: "Europe/London"
`,
			want: &Config{
				DefaultEntry: "log",
				Entries: []Entry{
					{
						ID: "log",
						Schedule: Schedule{
							Frequency: "daily",
							Days:      []int{1, 3, 5},
						},
						TemplateName: "log.tmpl",
					},
				},
				Paths: Paths{
					TemplatesDir:     "~/.journal/customtemplates",
					BaseDirectory:    "~/journals",
					JournalDirectory: "{{.Year}}/{{.Month}}/{{.Day}}/",
				},
				UserSettings: UserSettings{
					Timezone: "Europe/London",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid YAML Format",
			yamlData: `
defaultEntry: "task"
entries:
  - id: "task"
paths:
  baseDirectory: "journals"
  templatesDir
`,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempFile, err := os.CreateTemp("", "config-*.yaml")
			require.NoError(t, err)
			defer os.Remove(tempFile.Name())

			_, err = tempFile.WriteString(tt.yamlData)
			require.NoError(t, err)

			err = tempFile.Close()
			require.NoError(t, err)

			got, err := LoadConfig(tempFile.Name())

			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
