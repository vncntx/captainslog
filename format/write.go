package format

import (
	"os"
)

// Write a string to a stream
func Write(stream *os.File, str string) {
	_, _ = stream.WriteString(str)
}
