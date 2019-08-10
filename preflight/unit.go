package preflight

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// UnitTest provides utilities for unit testing
type UnitTest struct {
	*testing.T
}

// FileConsumer is a function that consumes a file
type FileConsumer func(file *os.File)

// Unit returns a new unit test
func Unit(t *testing.T) *UnitTest {
	return &UnitTest{t}
}

// Expect returns a new value-based expectation
func (unit *UnitTest) Expect(actual interface{}) Expectation {
	return ExpectValue(unit.T, actual)
}

// ExpectFile returns expectations based on file contents
func (unit *UnitTest) ExpectFile(file *os.File) Expectation {
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		unit.Errorf("could not read from file '%s'", file.Name())
	}
	return unit.Expect(string(contents))
}

// ExpectOutput returns a new output file-based expectation
func (unit *UnitTest) ExpectOutput(consumer FileConsumer) Expectation {
	readable, writable, err := os.Pipe()
	if err != nil {
		unit.Error(err)
	}

	// invoke the consumer
	consumer(writable)

	// close the write stream
	err = writable.Close()
	unit.Expect(err).Is().Nil()

	return unit.ExpectFile(readable)
}

// ExpectLog creates a set of expectations from log text
func (unit *UnitTest) ExpectLog(text string) *LogMessage {
	text, color := colorless(unit.T, text)
	parts := strings.Split(text, " :: ")

	fields := ""
	message := parts[3]
	if len(parts) > 4 {
		fields = parts[3]
		message = parts[4]
	}

	return &LogMessage{
		Time:     ExpectValue(unit.T, parts[1]),
		Name:     ExpectValue(unit.T, parts[2]),
		Level:    ExpectValue(unit.T, strings.Trim(parts[0], " ")),
		Color:    ExpectValue(unit.T, color),
		Fields:   ExpectValue(unit.T, fields),
		Message:  ExpectValue(unit.T, message),
		FullText: ExpectValue(unit.T, text),
	}
}

// ExpectLogs creates a list of expectations from logs
func (unit *UnitTest) ExpectLogs(consumer FileConsumer) []*LogMessage {
	readable, writable, err := os.Pipe()
	if err != nil {
		unit.Error(err)
	}

	// invoke the consumer
	consumer(writable)

	// close the write stream
	err = writable.Close()
	unit.Expect(err).Is().Nil()

	contents, err := ioutil.ReadAll(readable)
	if err != nil {
		unit.Errorf("could not read from logs '%s'", readable.Name())
	}
	var expectations []*LogMessage
	for _, line := range strings.Split(string(contents), "\n") {
		if len(line) > 0 {
			expectations = append(expectations, unit.ExpectLog(line))
		}
	}
	return expectations
}