package application

import (
	"errors"
	"fmt"

	"github.com/matthewchivers/journal/pkg/templating"
)

// PreparePatternData prepares the pattern data for the application
// This must be called before parsing any patterns
func (app *App) PreparePatternData() error {
	if app.LaunchTime.IsZero() {
		return errors.New("launch time must be set before preparing pattern data")
	}
	templateModel, err := templating.PrepareTemplateData(app.LaunchTime)
	if err != nil {
		return fmt.Errorf("failed to prepare template data: %w", err)
	}

	app.TemplateData = &templateModel
	return nil
}

// SetTopic sets the topic for the entry
func (app *App) SetTopic(topic string) error {
	if app.TemplateData == nil {
		return errors.New("pattern data must be initialised before setting topic")
	}
	if topic != "" {
		app.TemplateData.Topic = topic
	} else {
		entry, err := app.GetTargetEntry()
		if err != nil {
			return err
		}
		app.TemplateData.Topic = entry.Topic
	}
	return nil
}
