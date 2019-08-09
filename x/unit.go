package x

import (
	"github.com/vincentfiestada/captainslog/x/expect"
	"testing"
)

// Unit provides utilities for unit testing
type Unit struct {
	*testing.T
}

// NewUnit returns a new unit test
func NewUnit(test *testing.T) *Unit {
	return &Unit{test}
}

// Expect returns a new value-based expectation
func (unit *Unit) Expect(actual interface{}) expect.Expectation {
	return &expect.Value{
		Actual: actual,
	}
}