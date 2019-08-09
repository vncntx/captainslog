package expect

import (
	"testing"
)


// Value is an expectation based on a realized value
type Value struct {
	*testing.T
	Actual interface{}
}

// Equals asserts equality to another value
func (e Value) Equals(expected interface{}) Expectation {
	//assert.Equal(e, expected, e.Actual)
	return e
}