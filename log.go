package captainslog

import (
	"time"

	"github.com/vincentfiestada/captainslog/caller"
	"github.com/vincentfiestada/captainslog/format"
)

// Log levels
const (
	LogLevelTrace = iota
	LogLevelDebug = iota
	LogLevelInfo  = iota
	LogLevelWarn  = iota
	LogLevelError = iota
	LogLevelFatal = iota
	LogLevelQuiet = iota
)

// printFunc is a function that formats and prints
type printFunc func(string, ...interface{})

// Level specifies the message types and severity
type Level uint8

// Logger is an object for logging
type Logger struct {
	Level         Level
	Name          string
	LogFormat     format.Factory
	TimeFormat    string
	HasColor      bool
	MaxNameLength int
}

// NewLogger returns a new logger with the specified minimum logging level
func NewLogger() *Logger {
	return &Logger{
		Level:         LogLevelDebug,
		TimeFormat:    "01-02-2006 15:04:05 MST",
		HasColor:      true,
		MaxNameLength: 15,
		LogFormat:     format.FactoryOf(format.Flat()),
	}
}

// SetTimeFormat sets the time format
func (log *Logger) SetTimeFormat(timeFormat string) {
	log.TimeFormat = timeFormat
}

// SetName overrides the caller name
func (log *Logger) SetName(name string) {
	log.Name = name
}

// SetLevel sets the logging level
func (log *Logger) SetLevel(level Level) {
	log.Level = level
}

// getName returns the name of the logger or its caller
func (log *Logger) getName() string {
	if len(log.Name) > 0 {
		return log.Name
	}
	// if the logger has no name, return the name of the caller
	return caller.Shorten(caller.GetName(4), log.MaxNameLength)
}

// createMessage returns a new message
func (log *Logger) createMessage() *Message {
	msg := &Message{
		sep:       " :: ",
		time:      time.Now().Format(log.TimeFormat),
		name:      log.getName(),
		format:    log.LogFormat(),
		threshold: log.Level,
	}
	if !log.HasColor {
		msg.colorize = printf
	}
	return msg
}

// Field adds a data field to the log
func (log *Logger) Field(name string, value interface{}) *Message {
	return log.createMessage().Field(name, value)
}

// Trace logs a message with level Trace
func (log *Logger) Trace(format string, args ...interface{}) {
	log.createMessage().Trace(format, args...)
}

// Debug logs a message with level Debug
func (log *Logger) Debug(format string, args ...interface{}) {
	log.createMessage().Debug(format, args...)
}

// Info logs a message with level Info
func (log *Logger) Info(format string, args ...interface{}) {
	log.createMessage().Info(format, args...)
}

// Warn logs a message with level Warn
func (log *Logger) Warn(format string, args ...interface{}) {
	log.createMessage().Warn(format, args...)
}

// Error logs a message with level Error
func (log *Logger) Error(format string, args ...interface{}) {
	log.createMessage().Error(format, args...)
}

// Exit logs an error and exits with the given code
func (log *Logger) Exit(code int, format string, args ...interface{}) {
	log.createMessage().Exit(code, format, args...)
}

// Fatal logs an error and exits with code 1
func (log *Logger) Fatal(format string, args ...interface{}) {
	log.createMessage().Fatal(format, args...)
}

// Panic logs an error and panics
func (log *Logger) Panic(format string, args ...interface{}) {
	log.createMessage().Panic(format, args...)
}
