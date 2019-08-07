package captainslog

import (
	"os"
	"time"

	"github.com/vincentfiestada/captainslog/caller"
	"github.com/vincentfiestada/captainslog/format"
)

// Log levels
const (
	LevelTrace int = iota
	LevelDebug int = iota
	LevelInfo  int = iota
	LevelWarn  int = iota
	LevelError int = iota
	LevelFatal int = iota
	LevelQuiet int = iota
)

// Defaults
const (
	ISO8601 = "01-02-2006 15:04:05 MST"
)

// printFunc is a function that formats and prints
type printFunc func(string, ...interface{})

// Logger is an object for logging
type Logger struct {
	Name          string
	Level         int
	HasColor      bool
	TimeFormat    string
	MaxNameLength int
	Stdout        *os.File
	Stderr        *os.File
	format        format.Factory
}

// NewLogger returns a new logger with the specified minimum logging level
func NewLogger() *Logger {
	return &Logger{
		HasColor:      true,
		Level:         LevelDebug,
		TimeFormat:    ISO8601,
		MaxNameLength: 15,
		Stdout:        os.Stdout,
		Stderr:        os.Stderr,
		format:        format.Flat,
	}
}

// name returns the name of the logger or of its caller
func (log *Logger) name() string {
	if len(log.Name) > 0 {
		return log.Name
	}
	// if the logger has no name, return the name of the caller
	return caller.Shorten(caller.GetName(4), log.MaxNameLength)
}

// message returns a new message
func (log *Logger) message() *Message {
	msg := &Message{
		sep:       " :: ",
		time:      time.Now().Format(log.TimeFormat),
		name:      log.name(),
		stdout:    log.Stdout,
		stderr:    log.Stderr,
		hasColor:  log.HasColor,
		threshold: log.Level,
		format:    log.format(),
	}
	return msg
}

// Field adds a data field to the log
func (log *Logger) Field(name string, value interface{}) *Message {
	return log.message().Field(name, value)
}

// Trace logs a message with level Trace
func (log *Logger) Trace(format string, args ...interface{}) {
	log.message().Trace(format, args...)
}

// Debug logs a message with level Debug
func (log *Logger) Debug(format string, args ...interface{}) {
	log.message().Debug(format, args...)
}

// Info logs a message with level Info
func (log *Logger) Info(format string, args ...interface{}) {
	log.message().Info(format, args...)
}

// Warn logs a message with level Warn
func (log *Logger) Warn(format string, args ...interface{}) {
	log.message().Warn(format, args...)
}

// Error logs a message with level Error
func (log *Logger) Error(format string, args ...interface{}) {
	log.message().Error(format, args...)
}

// Exit logs an error and exits with the given code
func (log *Logger) Exit(code int, format string, args ...interface{}) {
	log.message().Exit(code, format, args...)
}

// Fatal logs an error and exits with code 1
func (log *Logger) Fatal(format string, args ...interface{}) {
	log.message().Fatal(format, args...)
}

// Panic logs an error and panics
func (log *Logger) Panic(format string, args ...interface{}) {
	log.message().Panic(format, args...)
}
