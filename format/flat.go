package format

import (
	"fmt"
	"os"

	"github.com/vincentfiestada/captainslog/v2/msg"
)

// Flat formats a message as flat text
func Flat(msg *msg.Message) {
	stream, level, colorize := msg.Props()
	if !msg.HasColor {
		colorize = fmt.Sprintf
	}

	Write(stream, colorize("%6s", level))
	separate(stream)
	Write(stream, msg.Time)
	separate(stream)
	Write(stream, colorize(msg.Name))
	if len(msg.Data) > 0 {
		separate(stream)
		for i := 0; i < len(msg.Data)-1; i += 2 {
			if i > 0 {
				Write(stream, ", ")
			}
			Write(stream, fmt.Sprintf("%s=%#v", msg.Data[i], msg.Data[i+1]))
		}
	}
	separate(stream)
	Write(stream, msg.Text)
	Write(stream, "\n")
}

// separate prints out a separator between parts of the message
func separate(stream *os.File) {
	Write(stream, " :: ")
}
