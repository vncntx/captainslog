package captainslog

import (
	"os"
	"time"

	"github.com/vincentfiestada/captainslog/caller"
	"github.com/vincentfiestada/captainslog/format"
	"github.com/vincentfiestada/captainslog/levels"
	"github.com/vincentfiestada/captainslog/msg"
)

// Defaults
const (
	ISO8601 = "01-02-2006 15:04:05 MST"
)

// Logger is an object for logging
type Logger struct {
	Name       string
	Level      int
	HasColor   bool
	TimeFormat string
	NameCutoff int
	Stdout     *os.File
	Stderr     *os.File
	Format     msg.Printer
}

// NewLogger returns a new logger with the specified minimum logging level
func NewLogger() *Logger {
	return &Logger{
		HasColor:   true,
		Level:      levels.Debug,
		TimeFormat: ISO8601,
		NameCutoff: 15,
		Stdout:     os.Stdout,
		Stderr:     os.Stderr,
		Format:     format.Flat,
	}
}

// name returns the name of the logger or of its caller
func (log *Logger) name() string {
	if len(log.Name) > 0 {
		return log.Name
	}
	// if the logger has no name, return the name of the caller
	return caller.Shorten(caller.GetName(4), log.NameCutoff)
}

// message returns a new message
func (log *Logger) message() *msg.Message {
	msg := msg.MsgPool.Get().(*msg.Message)
	msg.Time = time.Now().Format(log.TimeFormat)
	msg.Name = log.name()
	msg.Stdout = log.Stdout
	msg.Stderr = log.Stderr
	msg.HasColor = log.HasColor
	msg.Threshold = log.Level
	msg.Format = log.Format
	msg.Data = []interface{}{}
	return msg
}

// I returns a single field that can be added to logs
func (log *Logger) I(name string, value interface{}) msg.Field {
	return msg.Field{name, value}
}

// Field starts a message with a data field
func (log *Logger) Field(name string, value interface{}) *msg.Message {
	return log.message().Field(name, value)
}

// Fields starts a message with multiple data fields
func (log *Logger) Fields(fields ...msg.Field) *msg.Message {
	return log.message().Fields(fields...)
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
