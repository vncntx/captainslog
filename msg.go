package captainslog

import (
	"fmt"
	"os"

	"github.com/vincentfiestada/captainslog/format"
)

// SprintFunc formats and returns a string
type SprintFunc func(string, ...interface{}) string

// Message is a log message that gets built in multiple steps
type Message struct {
	time      string
	name      string
	text      string
	sep       string
	level     uint8
	threshold uint8
	format    format.Format
	hasColor  bool
	stdout    *os.File
	stderr    *os.File
}

// SetLevel sets the priority level for the message
func (msg *Message) SetLevel(level uint8) {
	msg.level = level
}

// GetLevel returns the appropriate level and color
func (msg *Message) GetLevel() (stream *os.File, level string, color SprintFunc) {
	switch msg.level {
	case LogLevelTrace:
		return msg.stdout, "trace", purple
	case LogLevelDebug:
		return msg.stdout, "debug", green
	case LogLevelInfo:
		return msg.stdout, "info", blue
	case LogLevelWarn:
		return msg.stderr, "warn", yellow
	case LogLevelError:
		return msg.stderr, "error", red
	default:
		return msg.stderr, "fatal", red
	}
}

// print outputs the constructed log
func (msg *Message) print() {

	if msg.level < msg.threshold {
		return
	}

	stream, level, colorize := msg.GetLevel()
	if !msg.hasColor {
		colorize = fmt.Sprintf
	}

	stream.WriteString(colorize("%6s", level))
	msg.separate(stream)
	stream.WriteString(msg.time)
	msg.separate(stream)
	stream.WriteString(colorize(msg.name))
	msg.separate(stream)
	if !msg.format.IsEmpty() {
		stream.WriteString(msg.format.GetFields())
		msg.separate(stream)
	}
	stream.WriteString(msg.text)
	stream.WriteString("\n")
}

// separate prints out a separator string
func (msg *Message) separate(stream *os.File) {
	stream.WriteString(msg.sep)
}

// Field adds a data field to the log
func (msg *Message) Field(name string, value interface{}) *Message {
	msg.format.AddField(name, value)
	return msg
}

// Log outputs the message with the specified level
func (msg *Message) Log(level uint8, format string, args ...interface{}) {
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
