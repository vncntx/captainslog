package preflight

import (
	"regexp"
	"strings"
	"testing"
)

// LogExpectation is a set of expectations based on parsed logs
type LogExpectation struct {
	Time     Expectation
	Name     Expectation
	Level    Expectation
	Fields   Expectation
	Message  Expectation
	FullText Expectation
}

// ExpectLog creates a new LogExpectation
func ExpectLog(t *testing.T, text string) *LogExpectation {
	text = colorless(t, text)
	parts := strings.Split(text, " :: ")

	fields := ""
	message := parts[3]
	if len(parts) > 4 {
		fields = parts[3]
		message = parts[4]
	}

	return &LogExpectation{
		Time:     ExpectValue(t, parts[1]),
		Name:     ExpectValue(t, parts[2]),
		Level:    ExpectValue(t, strings.Trim(parts[0], " ")),
		Fields:   ExpectValue(t, fields),
		Message:  ExpectValue(t, message),
		FullText: ExpectValue(t, text),
	}
}

// colorless removes ANSI color codes from text and returns the color
func colorless(t *testing.T, text string) string {
	rxp, err := regexp.Compile(`\x1b\[[0-9;]*m`)
	if err != nil {
		t.Error(err)
	}

	return rxp.ReplaceAllString(text, "")
}
