package format

import (
	"fmt"

	"captainslog/v2/msg"
)

// JSON formats a message as JSON
func JSON(msg *msg.Message) {
	stream, level, _ := msg.Props()

	Write(stream, fmt.Sprintf(`{"level":"%s","time":"%s","from":"%s",`, level, msg.Time, msg.Name))
	if len(msg.Data) > 0 {
		Write(stream, `"fields":{`)
		for i := 0; i < len(msg.Data)-1; i += 2 {
			if i > 0 {
				Write(stream, ",")
			}
			Write(stream, fmt.Sprintf(`"%s":%#v`, msg.Data[i], msg.Data[i+1]))
		}
		Write(stream, "},")
	}
	Write(stream, fmt.Sprintf("\"message\":\"%s\"}\n", msg.Text))
}
