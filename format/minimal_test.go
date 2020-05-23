package format_test

import (
	"os"
	"testing"

	"github.com/vincentfiestada/captainslog/v2/format"
	"github.com/vincentfiestada/captainslog/v2/levels"
	"github.com/vincentfiestada/captainslog/v2/msg"
	"github.com/vincentfiestada/captainslog/v2/preflight"
)

func TestMinimal(test *testing.T) {
	t := preflight.Unit(test)

	t.ExpectOutput(func(stdout *os.File) {
		message := &msg.Message{
			Time:      "08-28-2019 12:32:24 PST",
			Name:      "captainslog",
			Text:      "starship enterprise",
			Level:     levels.Info,
			Threshold: levels.Info,
			Stdout:    stdout,
			Print:     format.Minimal,
			Data: []interface{}{
				"captain",
				"picard",
				"first officer",
				"riker",
			},
		}

		message.Print(message)

	}).Equals("  info: [captain=\"picard\", first officer=\"riker\"] starship enterprise\n")
}
