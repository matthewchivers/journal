package application

import (
	"errors"

	"github.com/matthewchivers/journal/pkg/editor"
)

// SetEditor sets the editor ID for the entry
// If editor is empty, the default editor is used
func (app *App) SetEditor(editorID string) error {
	if app.targetEntry == nil {
		return errors.New("entry must be set before setting editor ID")
	}
	if editorID != "" {
		app.Editor = editorID
	}
	app.Editor = app.Config.Editor
	if app.targetEntry.Editor != "" {
		app.Editor = app.targetEntry.Editor
	}

	switch app.Editor {
	case "":
		return errors.New("editor not set")
	case "vscode":
		vscEditor, err := editor.NewVSCodeEditor()
		if err != nil {
			return err
		}
		app.targetEditor = vscEditor
	default:
		return errors.New("editor not supported")
	}

	return nil
}

// GetEditor returns the editor for the entry
func (app *App) GetEditor() (editor.Editor, error) {
	if app.targetEditor == nil {
		return nil, errors.New("editor must be set before getting editor")
	}
	return app.targetEditor, nil
}
