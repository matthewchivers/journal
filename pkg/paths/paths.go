package paths

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/matthewchivers/journal/pkg/caltools"
)

// ParsePathTemplate creates a new path for a journal entry based on a path template
func ParsePathTemplate(nestedPathTemplate string, docTypeName string) (string, error) {
	data := struct {
		Year           string
		Month          string
		Day            string
		WeekdayName    string
		WeekdayNumber  string
		WeekCommencing string
		WeekNumber     string
		DocTypeName    string
	}{
		Year:           time.Now().Format("2006"),
		Month:          time.Now().Format("01"),
		Day:            time.Now().Format("02"),
		WeekdayName:    time.Now().Weekday().String(),
		WeekdayNumber:  string(rune(time.Now().Weekday())),
		WeekCommencing: caltools.WeekCommencing(time.Now()).Format("2006-01-02"),
		WeekNumber:     string(rune(caltools.WeekOfMonth(time.Now()))),
		DocTypeName:    docTypeName,
	}

	// WeekCommencing directories should nest in the same Year/Month as the commencing date)
	if strings.Contains(nestedPathTemplate, "{{.WeekCommencing}}") {
		wc := caltools.WeekCommencing(time.Now())
		data.Year = wc.Format("2006")
		data.Month = wc.Format("01")
		data.WeekNumber = string(rune(caltools.WeekOfMonth(wc)))
	}

	t, err := template.New("path").Parse(nestedPathTemplate)
	if err != nil {
		return "", err
	}

	var path bytes.Buffer
	err = t.Execute(&path, data)
	if err != nil {
		return "", err
	}

	pathString := path.String()

	return pathString, nil
}

// ConstructFullPath constructs the full path for the new file
func ConstructFullPath(baseDir, nestedPathTemplate string, docTypeName string) (string, error) {
	nestedPath, err := ParsePathTemplate(nestedPathTemplate, docTypeName)
	if err != nil {
		return "", fmt.Errorf("failed to construct nested path: %w", err)
	}

	fileName := fmt.Sprintf("%s.md", docTypeName)
	fullPath := filepath.Join(baseDir, nestedPath, fileName)
	return fullPath, nil
}
