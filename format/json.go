package format

import (
	"fmt"

	"github.com/vincentfiestada/captainslog/v2/msg"
)

// JSON formats a message as JSON
func JSON(msg *msg.Message) {
	stream, level, _ := msg.Props()

	write(stream, fmt.Sprintf(`{"level":"%s","time":"%s","caller":"%s",`, level, msg.Time, msg.Name))
	if len(msg.Data) > 0 {
		for i := 0; i < len(msg.Data)-1; i += 2 {
			write(stream, fmt.Sprintf(`"%s":%#v,`, msg.Data[i], msg.Data[i+1]))
		}
	}
	write(stream, fmt.Sprintf("\"message\":\"%s\"}\n", msg.Text))
}
