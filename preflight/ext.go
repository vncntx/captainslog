// Extended preflight package
package preflight

import (
	"testing"

	"vincent.click/pkg/captainslog/v2/preflight/log"
	"vincent.click/pkg/preflight"
)

// Test is a preflight test with extensions
type Test struct {
	*preflight.Test
}

// Unit returns a new Test
func Unit(test *testing.T) *Test {
	return &Test{
		preflight.Unit(test),
	}
}

// ExpectLog returns a set of expectations from a log
func (t *Test) ExpectLog(text string) log.Expectations {
	return log.Expect(t.T, text)
}

// ExpectLogged returns expectations from a function that writes logs
func (t *Test) ExpectLogged(consumer log.LogsConsumer) (stdout []log.Expectations, stderr []log.Expectations) {
	return log.ExpectLogged(t.T, consumer)
}
