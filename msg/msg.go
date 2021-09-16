package msg

import (
	"fmt"
	"os"
	"sync"

	"vincent.click/pkg/captainslog/v2/levels"
	"vincent.click/pkg/preflight"
)

// Field is a key-value pair
type Field [2]interface{}

// Format formats and prints out a log message
type Format func(msg *Message)

// Message is a log message that gets built in multiple steps
type Message struct {
	Time      string
	Name      string
	Text      string
	Level     int
	Threshold int
	HasColor  bool
	Stdout    *os.File
	Stderr    *os.File
	Print     Format
	Data      []interface{}
}

// MsgPool is a synchronized pool of messages
var MsgPool = sync.Pool{
	New: func() interface{} {
		return &Message{}
	},
}

// Props returns the message stream, level, and color
func (msg *Message) Props() (stream *os.File, level string, color Color) {
	switch msg.Level {
	case levels.Trace:
		return msg.Stdout, "trace", cyan
	case levels.Debug:
		return msg.Stdout, "debug", green
	case levels.Info:
		return msg.Stdout, "info", blue
	case levels.Warn:
		return msg.Stderr, "warn", yellow
	case levels.Error:
		return msg.Stderr, "error", red
	default:
		return msg.Stderr, "fatal", red
	}
}

// Field adds a data field to the message
func (msg *Message) Field(name string, value interface{}) *Message {
	msg.Data = append(msg.Data, name, value)

	return msg
}

// Fields adds multiple fields to the message
func (msg *Message) Fields(fields ...Field) *Message {
	for _, field := range fields {
		msg.Data = append(msg.Data, field[0], field[1])
	}

	return msg
}

// Log outputs the message with the specified level
func (msg *Message) Log(level int, format string, args ...interface{}) {
	msg.Level = level
	if msg.Level < msg.Threshold {
		return
	}

	msg.Text = fmt.Sprintf(format, args...)
	msg.Print(msg)
	// Return message to pool
	MsgPool.Put(msg)
}

// Trace outputs the message with level Trace
func (msg *Message) Trace(format string, args ...interface{}) {
	msg.Log(levels.Trace, format, args...)
}

// Debug outputs the message with level Debug
func (msg *Message) Debug(format string, args ...interface{}) {
	msg.Log(levels.Debug, format, args...)
}

// Info outputs the message with level Info
func (msg *Message) Info(format string, args ...interface{}) {
	msg.Log(levels.Info, format, args...)
}

// Warn outputs the message with level Warn
func (msg *Message) Warn(format string, args ...interface{}) {
	msg.Log(levels.Warn, format, args...)
}

// Error outputs the message with level Error
func (msg *Message) Error(format string, args ...interface{}) {
	msg.Log(levels.Error, format, args...)
}

// Exit outputs the message as an error and exits with the given code
func (msg *Message) Exit(code int, format string, args ...interface{}) {
	msg.Log(levels.Fatal, format, args...)
	preflight.Captor.Exit(code)
}

// Fatal outputs the message as an error and exits with code 1
func (msg *Message) Fatal(format string, args ...interface{}) {
	msg.Log(levels.Fatal, format, args...)
	preflight.Captor.Exit(1)
}

// Panic outputs the message as an error and panics
func (msg *Message) Panic(format string, args ...interface{}) {
	msg.Log(levels.Fatal, format, args...)
	panic(fmt.Errorf(format, args...))
}
