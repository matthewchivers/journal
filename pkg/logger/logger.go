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
	loggerInstance *zerolog.Logger
	loggingPath    string
	logLevel       LogLevel
	logJSON        bool
)

// GetLogger returns the logger instance
// If the logger instance is not set, it will be initialized with the default log level
func GetLogger() (*zerolog.Logger, error) {
	if logLevel == notSet {
		logLevel = LogLevelDefault
	}
	if loggerInstance == nil {
		err := InitLogger(logLevel)
		if err != nil {
			return nil, err
		}
	}
	return loggerInstance, nil
}

// SetLogger sets the logger instance
func SetLogger(l *zerolog.Logger) {
	loggerInstance = l
}

// SetLoggingPath
func SetLoggingPath(logPath string) error {
	if logPath != "" {
		loggingPath = logPath
	} else {
		appHome, err := paths.GetAppHomePath()
		if err != nil {
			return err
		}
		loggingPath = filepath.Join(appHome, "journal.log")
	}
	return nil
}

// SetLogLevel sets the log level
func SetLogLevel(level LogLevel) {
	logLevel = level
}

// SetLogJSON sets the JSON flag
func SetLogJSON(json bool) {
	logJSON = json
}

// InitLogger initializes the logger
func InitLogger(lvl LogLevel) error {
	logLevel = lvl

	switch logLevel {
	case LogLevelDefault:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case LogLevelInfo:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case LogLevelDebug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	writers := []io.Writer{}

	cw := getConsoleWriter()
	if cw != nil {
		writers = append(writers, cw)
	}

	fw, err := getFileWriter()
	if err != nil {
		return fmt.Errorf("error getting file writer: %q", err)
	}
	writers = append(writers, fw)

	multiWriter := zerolog.MultiLevelWriter(writers...)
	multiLogger := zerolog.New(multiWriter).With().Timestamp().Logger()
	loggerInstance = &multiLogger

	loggerInstance.Debug().Str("log_level", loggerInstance.GetLevel().String()).
		Str("log_path", loggingPath).
		Msg("created logger")
	return nil
}

// getConsoleWriter returns the console writer
func getConsoleWriter() io.Writer {
	var iow io.Writer
	if logLevel >= LogLevelInfo {
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
	}
	return iow
}

func getFileWriter() (io.Writer, error) {
	// Structured logging to file
	if loggingPath == "" {
		if err := SetLoggingPath(""); err != nil {
			return nil, fmt.Errorf("error setting logging path: %q", err)
		}
	}

	file, err := os.OpenFile(loggingPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("error creating log file: %q with error %q", loggingPath, err)
	}
	fileWriter := zerolog.New(file).With().Timestamp().Logger().Output(file)
	return fileWriter, nil
}
