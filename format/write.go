package format

import (
	"io"
)

// Write a string to a stream
func Write(stream io.StringWriter, str string) {
	_, _ = stream.WriteString(str)
}
