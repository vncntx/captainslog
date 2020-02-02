package format_test

import (
	"os"
	"testing"

	"github.com/vincentfiestada/captainslog/v2/format"
	"github.com/vincentfiestada/captainslog/v2/levels"
	"github.com/vincentfiestada/captainslog/v2/msg"
	"github.com/vincentfiestada/captainslog/v2/preflight"
)

func TestJSON(test *testing.T) {
	t := preflight.Unit(test)

	t.ExpectOutput(func(stdout *os.File) {
		message := &msg.Message{
			Time:      "08-28-2019 12:32:24 PST",
			Name:      "captainslog",
			Text:      "starship enterprise",
			Level:     levels.Info,
			Threshold: levels.Info,
			Stdout:    stdout,
			Format:    format.JSON,
			Data: []interface{}{
				"captain",
				"picard",
				"first officer",
				"riker",
			},
		}

		format.JSON(message)

	}).Equals("{\"level\":\"info\",\"time\":\"08-28-2019 12:32:24 PST\",\"caller\":\"captainslog\",\"captain\":\"picard\",\"first officer\":\"riker\",\"message\":\"starship enterprise\"}\n")
}
