package format_test

import (
	"os"
	"testing"

	"github.com/vincentfiestada/captainslog/format"
	"github.com/vincentfiestada/captainslog/levels"
	"github.com/vincentfiestada/captainslog/msg"
	"github.com/vincentfiestada/captainslog/preflight"
)

func TestFlat(test *testing.T) {
	t := preflight.Unit(test)

	t.ExpectOutput(func(stdout *os.File) {
		message := &msg.Message{
			Time:      "08-28-2019 12:32:24 PST",
			Name:      "captainslog",
			Text:      "starship enterprise",
			Level:     levels.Info,
			Threshold: levels.Info,
			Stdout:    stdout,
			Format:    format.Flat,
			Fields: []interface{}{
				"captain",
				"picard",
				"first officer",
				"riker",
			},
		}

		format.Flat(message)

	}).Equals("  info :: 08-28-2019 12:32:24 PST :: captainslog :: captain=\"picard\", first officer=\"riker\" :: starship enterprise\n")
}
