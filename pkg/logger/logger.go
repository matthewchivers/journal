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
	LogLevelDefault LogLevel = iota
	LogLevelInfo    LogLevel = iota
	LogLevelDebug   LogLevel = iota
)

var (
	loggerInstance *zerolog.Logger
	loggingPath    string
	logLevel       LogLevel
)

// GetLogger returns the logger instance
func GetLogger() *zerolog.Logger {
	if loggerInstance == nil {
		InitLogger()
	}
	return loggerInstance
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

// InitLogger initializes the logger
func InitLogger() {
	switch logLevel {
	case LogLevelDefault:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case LogLevelInfo:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case LogLevelDebug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	writers := []io.Writer{}

	if logLevel >= LogLevelInfo {
		// Human readable console logger
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		writers = append(writers, consoleWriter)
	}

	// Structured logging to file
	file, err := os.OpenFile(loggingPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("error creating log file: %q\n", loggingPath)
		fmt.Println("failed to create log file:", err)
	}
	fileWriter := zerolog.New(file).With().Timestamp().Logger().Output(file)

	writers = append(writers, fileWriter)

	multiWriter := zerolog.MultiLevelWriter(writers...)
	multiLogger := zerolog.New(multiWriter).With().Timestamp().Logger()
	loggerInstance = &multiLogger

	loggerInstance.Debug().Str("log_level", loggerInstance.GetLevel().String()).
		Str("log_path", loggingPath).
		Msg("created logger")
}
