package format

import (
	"os"
)

// Write a string to a file
func Write(stream *os.File, str string) {
	_, _ = stream.WriteString(str)
}
