package format_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vincentfiestada/captainslog/format"
)

func TestFlat(t *testing.T) {
	// should return a new instance
	assert.NotNil(t, format.Flat())
}

func TestFlatGetFields(t *testing.T) {
	flat := format.Flat()
	flat.AddField("finn", "human")

	// should append a field to the string
	assert.Equal(t, "finn=\"human\"", flat.GetFields())

	flat.AddField("jake", "dog")
	// should separate fields with a space
	assert.Equal(t, "finn=\"human\" jake=\"dog\"", flat.GetFields())
}

func TestFlatIsEmpty(t *testing.T) {
	flat := format.Flat()

	// should be true if no fields have been added
	assert.True(t, flat.IsEmpty())

	flat.AddField("bubblegum", "candy")
	// should be false if one or more fields have been added
	assert.False(t, flat.IsEmpty())
}
