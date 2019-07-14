package format

import "fmt"

// FlatFormat specifies a simple, flat list of fields
type FlatFormat struct {
	fields string
}

// Flat returns a new flat format specifier
func Flat() *FlatFormat {
	return &FlatFormat{}
}

// AddField appends a field into the list
func (format *FlatFormat) AddField(name string, value interface{}) {
	if format.fields != "" {
		format.fields += " "
	}
	format.fields += fmt.Sprintf("%s=%#v", name, value)
}

// GetFields returns fields as a flat list
func (format *FlatFormat) GetFields() string {
	return format.fields
}

// IsEmpty returns true if no fields have been added
func (format *FlatFormat) IsEmpty() bool {
	return format.fields == ""
}
