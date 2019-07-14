package format

import (
	"encoding/json"
)

// JSONFormat specifies a JSON format
type JSONFormat struct {
	fields map[string]interface{}
}

// JSON returns a new JSON format specifier
func JSON() *JSONFormat {
	return &JSONFormat{
		fields: make(map[string]interface{}),
	}
}

// AddField sets a data field
func (format *JSONFormat) AddField(name string, value interface{}) {
	format.fields[name] = value
}

// GetFields returns fields as a JSON string
func (format *JSONFormat) GetFields() string {
	result, err := json.Marshal(format.fields)
	if err != nil {
		panic(err)
	}
	return string(result)
}

// IsEmpty returns true if no fields have been set
func (format *JSONFormat) IsEmpty() bool {
	return len(format.fields) == 0
}
