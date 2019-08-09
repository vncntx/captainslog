package expect

// Expectation represents an expectation about a value or behavior
type Expectation interface {
	Equals(expected interface{}) Expectation
}
