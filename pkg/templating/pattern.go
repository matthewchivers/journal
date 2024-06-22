package templating

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/matthewchivers/journal/pkg/config/app"
)

// ParsePattern creates a new path for a journal entry based on a path template
func ParsePattern(rawTemplate string, fileType app.FileType) (string, error) {
	data, err := PrepareTemplateData(fileType, strings.Contains(rawTemplate, "{{.WeekCommencing}}"))
	if err != nil {
		return "", err
	}

	t, err := template.New("path").Parse(rawTemplate)
	if err != nil {
		return "", err
	}

	var templateB bytes.Buffer
	err = t.Execute(&templateB, data)
	if err != nil {
		return "", err
	}

	parsedTemplate := templateB.String()

	return parsedTemplate, nil
}
