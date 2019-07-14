package captainslog

import (
	"fmt"
	"os"

	"github.com/vincentfiestada/captainslog/format"
)

// Message is a log message that gets built in multiple steps
type Message struct {
	time      string
	name      string
	text      string
	sep       string
	level     Level
	threshold Level
	format    format.Format
	colorize  func(string, ...interface{})
}

// SetLevel sets the priority level for the message
func (msg *Message) SetLevel(level Level) {
	msg.level = level
	if msg.colorize != nil {
		return
	}
	switch msg.level {
	case LogLevelTrace:
		msg.colorize = purple
		break
	case LogLevelDebug:
		msg.colorize = green
		break
	case LogLevelInfo:
		msg.colorize = blue
		break
	case LogLevelWarn:
		msg.colorize = yellow
		break
	case LogLevelError:
		msg.colorize = red
		break
	case LogLevelFatal:
		msg.colorize = red
		break
	default:
		msg.colorize = printf
	}
}

// GetLevel returns the message's priority level name
func (msg *Message) GetLevel() string {
	switch msg.level {
	case LogLevelTrace:
		return "trace"
	case LogLevelDebug:
		return "debug"
	case LogLevelInfo:
		return "info"
	case LogLevelWarn:
		return "warn"
	case LogLevelError:
		return "error"
	default:
		return "fatal"
	}
}

// print outputs the constructed log
func (msg *Message) print() {

	if msg.level < msg.threshold {
		return
	}

	msg.colorize("%6s", msg.GetLevel())
	msg.separate()
	fmt.Print(msg.time)
	msg.separate()
	msg.colorize(msg.name)
	msg.separate()
	if !msg.format.IsEmpty() {
		fmt.Print(msg.format.GetFields())
		msg.separate()
	}
	fmt.Println(msg.text)
}

// separate prints out a separator string
func (msg *Message) separate() {
	fmt.Print(msg.sep)
}

// Field adds a data field to the log
func (msg *Message) Field(name string, value interface{}) *Message {
	msg.format.AddField(name, value)
	return msg
}

// Log outputs the message with the specified level
func (msg *Message) Log(level Level, format string, args ...interface{}) {
	msg.SetLevel(level)
	msg.text = fmt.Sprintf(format, args...)
	msg.print()
}

// Trace outputs the message with level Trace
func (msg *Message) Trace(format string, args ...interface{}) {
	msg.Log(LogLevelTrace, format, args...)
}

// Debug outputs the message with level Debug
func (msg *Message) Debug(format string, args ...interface{}) {
	msg.Log(LogLevelDebug, format, args...)
}

// Info outputs the message with level Info
func (msg *Message) Info(format string, args ...interface{}) {
	msg.Log(LogLevelInfo, format, args...)
}

// Warn outputs the message with level Warn
func (msg *Message) Warn(format string, args ...interface{}) {
	msg.Log(LogLevelWarn, format, args...)
}

// Error outputs the message with level Error
func (msg *Message) Error(format string, args ...interface{}) {
	msg.Log(LogLevelError, format, args...)
}

// Exit outputs the message as an error and exits with the given code
func (msg *Message) Exit(code int, format string, args ...interface{}) {
	msg.Log(LogLevelFatal, format, args...)
	os.Exit(code)
}

// Fatal outputs the message as an error and exits with code 1
func (msg *Message) Fatal(format string, args ...interface{}) {
	msg.Log(LogLevelFatal, format, args...)
	os.Exit(1)
}

// Panic outputs the message as an error and panics
func (msg *Message) Panic(format string, args ...interface{}) {
	msg.Log(LogLevelFatal, format, args...)
	panic(fmt.Errorf(format, args...))
}
