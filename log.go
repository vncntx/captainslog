package captainslog

import (
	"fmt"
	"time"

	"github.com/vincentfiestada/captainslog/caller"

	"github.com/fatih/color"
)

// Log levels
const (
	LogLevelSilly   = iota
	LogLevelDebug   = iota
	LogLevelVerbose = iota
	LogLevelInfo    = iota
	LogLevelWarn    = iota
	LogLevelError   = iota
	LogLevelQuiet   = iota
)

// color print functions
var (
	pink   = color.New(color.FgMagenta).PrintfFunc()
	cyan   = color.New(color.FgCyan).PrintfFunc()
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
	LogLevel   Level
	TimeFormat string
	Name       string
	HasColor   bool

	callerDepth int
}

// NewLogger returns a new logger with the specified minimum logging level
func NewLogger() *Logger {
	return &Logger{
		LogLevel:    LogLevelDebug,
		TimeFormat:  "01-02-2006 15:04:05 MST",
		HasColor:    true,
		callerDepth: 4,
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
	log.LogLevel = level
}

// Log messages with the specified level
func (log *Logger) Log(level Level, format string, args ...interface{}) {
	if level < log.LogLevel {
		return
	}
	var printer printFunc
	var levelName string
	switch level {
	case LogLevelSilly:
		printer = pink
		levelName = "silly"
		break
	case LogLevelDebug:
		printer = cyan
		levelName = "debug"
		break
	case LogLevelVerbose:
		printer = green
		levelName = "verbose"
		break
	case LogLevelInfo:
		printer = blue
		levelName = "info"
		break
	case LogLevelWarn:
		printer = yellow
		levelName = "warn"
	default:
		printer = red
		levelName = "error"
	}
	if !log.HasColor {
		printer = printf
	}
	printer("%7s", levelName)
	fmt.Printf(":: %s :: ", time.Now().Format(log.TimeFormat))
	printer("%s :: ", log.getName())
	fmt.Printf(format, args...)
	fmt.Println()
}

// getName returns the name of the logger or its caller
func (log *Logger) getName() string {
	if len(log.Name) > 0 {
		return log.Name
	}
	// if the logger has no name, return the name of the caller
	return caller.GetName(log.callerDepth)
}

// Silly logs messages with level Silly
func (log *Logger) Silly(format string, args ...interface{}) {
	log.Log(LogLevelSilly, format, args...)
}

// Debug logs messages with level Debug
func (log *Logger) Debug(format string, args ...interface{}) {
	log.Log(LogLevelDebug, format, args...)
}

// Verbose logs messages with level Verbose
func (log *Logger) Verbose(format string, args ...interface{}) {
	log.Log(LogLevelVerbose, format, args...)
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

// printf a simple colorless print function
func printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
