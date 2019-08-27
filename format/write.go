package format

import (
	"log"
	"os"
)

// write a string to a file
func write(stream *os.File, str string) {
	_, err := stream.WriteString(str)
	if err != nil {
		log.Println("captainslog error!", err)
	}
}
