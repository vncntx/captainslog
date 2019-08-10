package preflight

import (
	"regexp"
	"testing"
)

// LogMessage is a set of expectations based on parsed logs
type LogMessage struct {
	Time     Expectation
	Name     Expectation
	Level    Expectation
	Color    Expectation
	Fields   Expectation
	Message  Expectation
	FullText Expectation
}

// colorless removes ANSI color codes from text and returns the color
func colorless(t *testing.T, text string) (colorless string, color string) {
	rxp, err := regexp.Compile("\\x1b\\[[0-9;]*m")
	if err != nil {
		t.Error(err)
	}
	colorCode := rxp.FindString(text)
	if len(colorCode) >= 3 {
		// get the color code if it is present
		color = colorCode[len(colorCode)-3:]
	}
	return rxp.ReplaceAllString(text, ""), color
}
