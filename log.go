package captainslog

import (
	"fmt"
	"os"
	"time"

	"github.com/vincentfiestada/captainslog/caller"

	"github.com/fatih/color"
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

// Error codes
const errFatal = 1

// color print functions
var (
	pink   = color.New(color.FgMagenta).PrintfFunc()
	blue   = color.New(color.FgBlue).PrintfFunc()
	green  = color.New(color.FgGreen).PrintfFunc()
	yellow = color.New(color.FgYellow).PrintfFunc()
	red    = color.New(color.FgRed).PrintfFunc()
)

// printFunc is a function that formats and prints
type printFunc func(string, ...interface{})

// Level specifies the message types and severity
type Level uint8

// Logger is an object for logging
type Logger struct {
	Level         Level
	TimeFormat    string
	Name          string
	HasColor      bool
	MaxNameLength int

	callerDepth int
}

// NewLogger returns a new logger with the specified minimum logging level
func NewLogger() *Logger {
	return &Logger{
		Level:         LogLevelDebug,
		TimeFormat:    "01-02-2006 15:04:05 MST",
		HasColor:      true,
		MaxNameLength: 15,
		callerDepth:   4,
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

// Log messages with the specified level
func (log *Logger) Log(level Level, format string, args ...interface{}) {
	if level < log.Level {
		return
	}
	var printer printFunc
	var levelName string
	switch level {
	case LogLevelTrace:
		printer = pink
		levelName = "trace"
		break
	case LogLevelDebug:
		printer = green
		levelName = "debug"
		break
	case LogLevelInfo:
		printer = blue
		levelName = "info"
		break
	case LogLevelWarn:
		printer = yellow
		levelName = "warn"
	case LogLevelError:
		printer = red
		levelName = "error"
	default:
		printer = red
		levelName = "fatal"
	}
	if !log.HasColor {
		printer = printf
	}
	printer("%7s", levelName)
	printf(" :: %s :: ", time.Now().Format(log.TimeFormat))
	printer("%s", log.getName())
	printf(" :: ")
	printf(format, args...)
	fmt.Println()
}

// getName returns the name of the logger or its caller
func (log *Logger) getName() string {
	if len(log.Name) > 0 {
		return log.Name
	}
	// if the logger has no name, return the name of the caller
	return caller.Shorten(caller.GetName(log.callerDepth), log.MaxNameLength)
}

// Trace logs messages with level Trace
func (log *Logger) Trace(format string, args ...interface{}) {
	log.Log(LogLevelTrace, format, args...)
}

// Debug logs messages with level Debug
func (log *Logger) Debug(format string, args ...interface{}) {
	log.Log(LogLevelDebug, format, args...)
}

// Info logs messages with level Info
func (log *Logger) Info(format string, args ...interface{}) {
	log.Log(LogLevelInfo, format, args...)
}

// Warn logs messages with level Warn
func (log *Logger) Warn(format string, args ...interface{}) {
	log.Log(LogLevelWarn, format, args...)
}

// Error logs messages with level Error
func (log *Logger) Error(format string, args ...interface{}) {
	log.Log(LogLevelError, format, args...)
}

// Exit logs an error and exits with an error code
func (log *Logger) Exit(code int, format string, args ...interface{}) {
	log.Log(LogLevelFatal, format, args...)
	os.Exit(code)
}

// Fatal logs an error and exits with error code 1
func (log *Logger) Fatal(format string, args ...interface{}) {
	log.Log(LogLevelFatal, format, args...)
	os.Exit(errFatal)
}

// Panic log an error and panics
func (log *Logger) Panic(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Log(LogLevelFatal, msg)
	panic(msg)
}

// printf a simple colorless print function
func printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
