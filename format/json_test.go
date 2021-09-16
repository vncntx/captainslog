package format_test

import (
	"os"
	"testing"

	"vincent.click/pkg/captainslog/v2/format"
	"vincent.click/pkg/captainslog/v2/levels"
	"vincent.click/pkg/captainslog/v2/msg"
	"vincent.click/pkg/preflight"
)

func TestJSON(test *testing.T) {
	t := preflight.Unit(test)

	w := t.ExpectWritten(func(stdout *os.File) {

		message := &msg.Message{
			Time:      "08-28-2019 12:32:24 PST",
			Name:      "captainslog",
			Text:      "starship enterprise",
			Level:     levels.Info,
			Threshold: levels.Info,
			Stdout:    stdout,
			Print:     format.JSON,
			Data: []interface{}{
				"captain",
				"picard",
				"first officer",
				"riker",
			},
		}

		message.Print(message)

	})
	defer w.Close()

	w.Text().Equals("{\"level\":\"info\",\"time\":\"08-28-2019 12:32:24 PST\",\"from\":\"captainslog\",\"fields\":{\"captain\":\"picard\",\"first officer\":\"riker\"},\"message\":\"starship enterprise\"}\n")
}
