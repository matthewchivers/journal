package templating

import (
	"bytes"
	"html/template"
	"strings"
	"time"

	"github.com/matthewchivers/journal/pkg/caltools"
	"github.com/matthewchivers/journal/pkg/config"
)

// ParsePattern creates a new path for a journal entry based on a path template
func ParsePattern(rawTemplate string, fileType config.FileType) (string, error) {
	timeNow := time.Now()
	data := struct {
		Year           string
		Month          string
		Day            string
		WeekdayName    string
		WeekdayNumber  string
		WeekCommencing string
		WeekNumber     string
		FileTypeName   string
		FileExtension  string
	}{
		Year:           timeNow.Format("2006"),
		Month:          timeNow.Format("01"),
		Day:            timeNow.Format("02"),
		WeekdayName:    timeNow.Weekday().String(),
		WeekdayNumber:  string(rune(timeNow.Weekday())),
		WeekCommencing: caltools.WeekCommencing(timeNow).Format("2006-01-02"),
		WeekNumber:     string(rune(caltools.WeekOfMonth(timeNow))),
		FileTypeName:   fileType.Name,
		FileExtension:  fileType.FileExtension,
	}

	// WeekCommencing directories should nest in the same Year/Month as the commencing date)
	if strings.Contains(rawTemplate, "{{.WeekCommencing}}") {
		wc := caltools.WeekCommencing(timeNow)
		data.Year = wc.Format("2006")
		data.Month = wc.Format("01")
		data.WeekNumber = string(rune(caltools.WeekOfMonth(wc)))
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
