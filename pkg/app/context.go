package app

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/matthewchivers/journal/pkg/config"
	"github.com/matthewchivers/journal/pkg/templating"
)

type Context struct {
	// LaunchTime is the time the application was launched
	LaunchTime time.Time

	// ConfigPath is the path to the configuration file
	ConfigPath string

	// Config is the application configuration
	Config *config.Config

	// EntryID is the ID of the entry
	EntryID string

	// Directory is the directory in which to create the entry
	Directory string

	// FileName is the name of the file to create
	FileName string

	// FilePath is the full path to the file
	FilePath string

	// TemplateData is the data used to populate the templating patterns
	TemplateData *templating.TemplateModel

	// entry is the entry configuration (used for convenience)
	entry *config.Entry
}

// SetLaunchTime sets the launch time of the application
func (ctx *Context) SetLaunchTime(launchTime time.Time) {
	ctx.LaunchTime = launchTime
}

// SetConfigPath sets the path to the configuration file
func (ctx *Context) SetConfigPath(cfgPathOverride string) error {
	if cfgPathOverride != "" {
		ctx.ConfigPath = cfgPathOverride
	} else {
		defaultConfigPath, err := config.GetDefaultConfigPath()
		if err != nil {
			return err
		}
		ctx.ConfigPath = defaultConfigPath
	}
	return nil
}

// SetupConfig loads the configuration from the specified path or the default path if specified path is empty
func (ctx *Context) SetupConfig() error {
	// Load the configuration
	cfg, err := config.LoadConfig(ctx.ConfigPath)
	if err != nil {
		return err
	}
	if err := cfg.Validate(); err != nil {
		return err
	}
	ctx.Config = cfg
	return nil
}

// SetEntryID sets the entry ID for the context
// If entryID is empty, the default entry is used
func (ctx *Context) SetEntryID(entryID string) error {
	if ctx.Config == nil {
		return errors.New("config must be loaded before setting entry ID")
	}
	if ctx.TemplateData == nil {
		return errors.New("pattern data must be initialised before setting entry id")
	}
	if entryID != "" {
		ctx.EntryID = strings.ToLower(entryID)
	} else {
		ctx.EntryID = strings.ToLower(ctx.Config.DefaultEntry)
	}
	if ctx.EntryID == "" {
		return errors.New("no entry specified")
	}
	if _, err := ctx.Config.GetEntry(ctx.EntryID); err != nil {
		return err
	}

	ctx.TemplateData.EntryID = ctx.EntryID
	return nil
}

func (ctx *Context) SetDirectory(dir string) error {
	if dir == "" {
		entryPath, err := ctx.GetEntryDir()
		if err != nil {
			return err
		}
		ctx.Directory = entryPath
	} else {
		ctx.Directory = dir
	}
	return nil
}

// SetFileName sets the file name for the entry
// If fileName is empty, the default file name is retrieved
func (ctx *Context) SetFileName(fileName string) error {
	if ctx.TemplateData == nil {
		return errors.New("pattern data must be initialised before setting file name")
	}
	if fileName != "" {
		ctx.FileName = fileName
	} else {
		fileName, err := ctx.GetEntryFileName()
		if err != nil {
			return err
		}
		ctx.FileName = fileName
	}
	return nil
}

func (ctx *Context) GetEntryDir() (string, error) {
	entry, err := ctx.GetEntry()
	if err != nil {
		return "", err
	}

	journalDirPattern := ctx.Config.Paths.JournalDirectory
	if entry.JournalDirOverride != "" {
		journalDirPattern = entry.JournalDirOverride
	}

	if ctx.TemplateData == nil {
		return "", errors.New("template data must be initialised before getting entry directory")
	}

	journalPath, err := ctx.TemplateData.ParsePattern(journalDirPattern)
	if err != nil {
		return "", fmt.Errorf("failed to construct journal path: %w", err)
	}

	nestedPath, err := ctx.TemplateData.ParsePattern(entry.Directory)
	if err != nil {
		return "", fmt.Errorf("failed to construct nested path: %w", err)
	}

	fullPath := filepath.Join(ctx.Config.Paths.BaseDirectory, journalPath, nestedPath)

	return fullPath, nil
}

// GetEntryFileName returns the file name for the entry
func (ctx *Context) GetEntryFileName() (string, error) {
	entry, err := ctx.GetEntry()
	if err != nil {
		return "", err
	}

	if ctx.TemplateData == nil {
		return "", errors.New("pattern data must be initialised before getting entry file name")
	}

	fileName, err := ctx.TemplateData.ParsePattern(entry.FileName)
	if err != nil {
		return "", fmt.Errorf("failed to construct file name: %w", err)
	}

	return fileName, nil
}

// GetFilePath returns the full path to the file
func (ctx *Context) GetFilePath() (string, error) {
	if ctx.FilePath == "" {
		if ctx.Directory == "" {
			return "", errors.New("directory must be set before getting file path")
		}
		if ctx.FileName == "" {
			return "", errors.New("file name must be set before getting file path")
		}
		ctx.FilePath = filepath.Join(ctx.Directory, ctx.FileName)
	}
	return ctx.FilePath, nil
}

func (ctx *Context) PreparePatternData() error {
	if ctx.LaunchTime.IsZero() {
		return errors.New("launch time must be set before preparing pattern data")
	}
	templateModel, err := templating.PrepareTemplateData(ctx.LaunchTime)
	if err != nil {
		return fmt.Errorf("failed to prepare template data: %w", err)
	}
	ctx.TemplateData = &templateModel
	return nil
}

// SetFileExtension sets the file extension for the entry
func (ctx *Context) SetFileExtension(fileExtension string) error {
	if ctx.TemplateData == nil {
		return errors.New("template data must be initialised before setting file extension")
	}
	if fileExtension != "" {
		ctx.TemplateData.FileExtension = fileExtension
	} else {
		entry, err := ctx.Config.GetEntry(ctx.EntryID)
		if err != nil {
			return err
		}
		if entry.FileExtension != "" {
			ctx.TemplateData.FileExtension = entry.FileExtension
		} else {
			ctx.TemplateData.FileExtension = ctx.Config.DefaultFileExtension
		}
	}
	return nil
}

// SetTopic sets the topic for the entry
func (ctx *Context) SetTopic(topic string) error {
	if ctx.TemplateData == nil {
		return errors.New("pattern data must be initialised before setting topic")
	}
	if topic != "" {
		ctx.TemplateData.Topic = topic
	} else {
		entry, err := ctx.GetEntry()
		if err != nil {
			return err
		}
		ctx.TemplateData.Topic = entry.Topic
	}
	return nil
}

func (ctx *Context) GetEntry() (*config.Entry, error) {
	if ctx.entry != nil {
		return ctx.entry, nil
	}
	if ctx.Config == nil {
		return nil, errors.New("config must be loaded before getting entry")
	}
	if ctx.EntryID == "" {
		return nil, errors.New("entry ID must be set before getting entry")
	}
	ent, err := ctx.Config.GetEntry(ctx.EntryID)
	if err != nil {
		return nil, err
	}
	ctx.entry = ent
	return ent, nil
}
