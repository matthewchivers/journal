package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
defaultDocType: "report"
documentTypes:
  - name: "report"
  paths:
	journalBaseDir: "journals"
  	documentNestingPath: "templates/report.tmpl"
`,
			want: &Config{
				DefaultDocType: "report",
				DocumentTypes: []DocumentType{
					{
						Name: "report",
					},
				},
				Paths: Paths{
					JournalBaseDir:     "journals",
					NestedPathTemplate: "templates/report.tmpl",
				},
			},
			wantErr: false,
		},
		{
			name: "Full Configuration With Schedules",
			yamlData: `
defaultDocType: "log"
documentTypes:
  - name: "log"
    schedule:
        frequency: "weekly"
        days: [1,3,5]
paths:
	templatesDir: "templates"
	journalBaseDir: "journals"
	nestedPathTemplate: "templates/log.tmpl"
userSettings:
    timezone: "Europe/London"
`,
			want: &Config{
				DefaultDocType: "log",
				DocumentTypes: []DocumentType{
					{
						Name: "log",
						Schedule: Schedule{
							Frequency: "weekly",
							Days:      []int{1, 3, 5},
						},
					},
				},
				Paths: Paths{
					TemplatesDir:       "templates",
					JournalBaseDir:     "journals",
					NestedPathTemplate: "templates/log.tmpl",
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
		defaultDocType: "task"
		documentTypes:
		  - name: "task"
		    templatePath: "templates/task.tmpl"
		paths:
		  journalDir: "journals"
		  templatesDir
		`,
			want:    nil,
			wantErr: true,
		},
		{
			name: "Missing Required Fields",
			yamlData: `
		documentTypes:
		  - name: "note"
		    templatePath: "templates/note.tmpl"
		`,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempFile, err := os.CreateTemp("", "config-*.yaml")
			if err != nil {
				t.Fatalf("Failed to create temporary config file: %v", err)
			}
			defer os.Remove(tempFile.Name())

			if _, err := tempFile.WriteString(tt.yamlData); err != nil {
				t.Fatalf("Failed to write to temporary config file: %v", err)
			}

			got, err := LoadConfig(tempFile.Name())

			if tt.wantErr {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
				if err != nil {
					return
				}
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
