package format

import (
	"fmt"
	"os"

	"github.com/vincentfiestada/captainslog/msg"
)

// Flat formats a message as flat text
func Flat(msg *msg.Message) {
	stream, level, colorize := msg.Props()
	if !msg.HasColor {
		colorize = fmt.Sprintf
	}

	write(stream, colorize("%6s", level))
	separate(stream)
	write(stream, msg.Time)
	separate(stream)
	write(stream, colorize(msg.Name))
	if len(msg.Fields) > 0 {
		separate(stream)
		for i := 0; i < len(msg.Fields)-1; i += 2 {
			if i > 0 {
				write(stream, ", ")
			}
			write(stream, fmt.Sprintf("%s=%#v", msg.Fields[i], msg.Fields[i+1]))
		}
	}
	separate(stream)
	write(stream, msg.Text)
	write(stream, "\n")
}

// separate prints out a separator between parts of the message
func separate(stream *os.File) {
	write(stream, " :: ")
}
