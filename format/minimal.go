package format

import (
	"fmt"

	"captainslog/v2/msg"
)

// Minimal prints a minimal log with no timestamp or name
func Minimal(msg *msg.Message) {
	stream, level, colorize := msg.Props()
	if !msg.HasColor {
		colorize = fmt.Sprintf
	}

	Write(stream, colorize("%6s", level))
	Write(stream, ": ")
	if len(msg.Data) > 0 {
		Write(stream, "[")
		for i := 0; i < len(msg.Data)-1; i += 2 {
			if i > 0 {
				Write(stream, ", ")
			}
			Write(stream, fmt.Sprintf("%s=%#v", msg.Data[i], msg.Data[i+1]))
		}
		Write(stream, "] ")
	}
	Write(stream, msg.Text)
	Write(stream, "\n")
}
