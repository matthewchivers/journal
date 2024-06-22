package templating

import (
	"bytes"
	"html/template"
	"strings"
)

// ParsePattern creates a new path for a journal entry based on a path template
func ParsePattern(rawTemplate string, entryID string, fileExtension string) (string, error) {
	data, err := PrepareTemplateData(entryID, fileExtension, strings.Contains(rawTemplate, "{{.WeekCommencing}}"))
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
