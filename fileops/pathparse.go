package fileops

import (
	"bytes"
	"strings"
	"text/template"
	"time"

	"github.com/matthewchivers/journal/caltools"
)

// ConstructJournalPath creates a new path for a journal entry based on a path template
func ConstructPath(pathTemplate, docTempName string) (string, error) {
	data := struct {
		Year           string
		Month          string
		Day            string
		WeekdayName    string
		WeekdayNumber  string
		WeekCommencing string
		WeekNumber     string
		TemplateName   string
	}{
		Year:           time.Now().Format("2006"),
		Month:          time.Now().Format("01"),
		Day:            time.Now().Format("02"),
		WeekdayName:    time.Now().Weekday().String(),
		WeekdayNumber:  string(rune(time.Now().Weekday())),
		WeekCommencing: caltools.WeekCommencing(time.Now()).Format("2006-01-02"),
		WeekNumber:     string(rune(caltools.WeekOfMonth(time.Now()))),
		TemplateName:   docTempName,
	}

	// WeekCommencing directories should nest in the same Year/Month as the commencing date)
	if strings.Contains(pathTemplate, "{{.WeekCommencing}}") {
		wc := caltools.WeekCommencing(time.Now())
		data.Year = wc.Format("2006")
		data.Month = wc.Format("01")
		data.WeekNumber = string(rune(caltools.WeekOfMonth(wc)))
	}

	t, err := template.New("path").Parse(pathTemplate)
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
