package logger

// Singleton logger instance

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/matthewchivers/journal/pkg/paths"
	"github.com/rs/zerolog"
)

// Struct with iota log levels
type LogLevel int

const (
	notSet          LogLevel = iota
	LogLevelDefault LogLevel = iota
	LogLevelInfo    LogLevel = iota
	LogLevelDebug   LogLevel = iota
)

var (
	loggingPath string
	logJSON     bool
	Log         *zerolog.Logger
)

// SetLogger sets the logger instance
func SetLogger(l *zerolog.Logger) {
	Log = l
}

// SetLogJSON sets the JSON flag
func SetLogJSON(json bool) {
	logJSON = json
}

// InitLogger initializes the logger
func InitLogger(level LogLevel) error {
	switch level {
	case LogLevelDefault:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case LogLevelInfo:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case LogLevelDebug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	writers := []io.Writer{}

	if level >= LogLevelInfo {
		cw := getConsoleWriter()
		if cw != nil {
			writers = append(writers, cw)
		}
	}

	fw, err := getFileWriter()
	if err != nil {
		return fmt.Errorf("error getting file writer: %q", err)
	}
	writers = append(writers, fw)

	multiWriter := zerolog.MultiLevelWriter(writers...)
	multiLogger := zerolog.New(multiWriter).With().Timestamp().Logger()
	Log = &multiLogger

	Log.Debug().Str("log_level", Log.GetLevel().String()).
		Str("log_path", loggingPath).
		Msg("created logger")
	return nil
}

// getConsoleWriter returns the console writer
func getConsoleWriter() io.Writer {
	var iow io.Writer
	if logJSON {
		iow = os.Stdout
	} else {
		// Human readable console logger
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		consoleWriter.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("*** %s ****", i)
		}
		iow = consoleWriter
	}
	return iow
}

func getFileWriter() (io.Writer, error) {
	// Structured logging to file
	appHome, err := paths.GetAppHomePath()
	if err != nil {
		return nil, err
	}
	loggingPath = filepath.Join(appHome, "journal.log")

	file, err := os.OpenFile(loggingPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("error creating log file: %q with error %q", loggingPath, err)
	}
	fileWriter := zerolog.New(file).With().Timestamp().Logger().Output(file)
	return fileWriter, nil
}
