package log

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"

	"vincent.click/pkg/preflight/expect"
)

// Expectations is a set of expectations about a log
type Expectations struct {
	Time     expect.Expectation
	Name     expect.Expectation
	Level    expect.Expectation
	Fields   expect.Expectation
	Message  expect.Expectation
	FullText expect.Expectation
}

// ExpectLog creates Expectations from a log
func Expect(t *testing.T, text string) Expectations {
	text = colorless(t, text)
	parts := strings.Split(text, " :: ")

	fields := ""
	message := parts[3]
	if len(parts) > 4 {
		fields = parts[3]
		message = parts[4]
	}

	return Expectations{
		Time:     expect.Value(t, parts[1]),
		Name:     expect.Value(t, parts[2]),
		Level:    expect.Value(t, strings.Trim(parts[0], " ")),
		Fields:   expect.Value(t, fields),
		Message:  expect.Value(t, message),
		FullText: expect.Value(t, text),
	}
}

// ExpectLogged creates expectations from a function that writes logs
func ExpectLogged(t *testing.T, consumer LogsConsumer) (stdout []Expectations, stderr []Expectations) {
	stdoutR, stdoutW := createPipe(t)
	stderrR, stderrW := createPipe(t)

	// invoke the consumer
	consumer(stdoutW, stderrW)

	// close the write streams
	closeStream(t, stdoutW)
	closeStream(t, stderrW)

	// create expectations by parsing the logs
	stdout = expectLogs(t, stdoutR)
	stderr = expectLogs(t, stderrR)

	return
}

// colorless removes ANSI color codes from text and returns the color
func colorless(t *testing.T, text string) string {
	rxp, err := regexp.Compile(`\x1b\[[0-9;]*m`)
	if err != nil {
		t.Error(err)
	}

	return rxp.ReplaceAllString(text, "")
}

// createPipe returns a connected set of new read and write streams
func createPipe(t *testing.T) (*os.File, *os.File) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("failed to create stream: %s", err)
	}

	return r, w
}

// closeStream closes a stream and handles errors
func closeStream(t *testing.T, f *os.File) {
	if err := f.Close(); err != nil {
		t.Errorf("failed to close stream: %s", err)
	}
}

// expectLogs returns expectations from multiple logs in a file
func expectLogs(t *testing.T, f *os.File) (expectations []Expectations) {

	contents, err := ioutil.ReadAll(f)
	if err != nil {
		t.Errorf("failed to read from stream: %s", err)

		return
	}

	for _, line := range strings.Split(string(contents), "\n") {
		if len(line) > 0 {
			expectations = append(expectations, Expect(t, line))
		}
	}

	return
}
