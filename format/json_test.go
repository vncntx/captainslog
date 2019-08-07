package format_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vincentfiestada/captainslog/format"
	"testing"
)

func TestJSON(t *testing.T) {
	// should return a new instance
	assert.NotNil(t, format.JSON())
}

func TestJSONGetFields(t *testing.T) {
	json := format.JSON()
	json.AddField("finn", "human")
	json.AddField("jake", "dog")

	// should add JSON fields
	assert.Equal(t, "{\"finn\":\"human\",\"jake\":\"dog\"}", json.GetFields())

	json.AddField("finn", "heroic")
	json.AddField("jake", "stretchy")

	// should replace existing fields
	assert.Equal(t, "{\"finn\":\"heroic\",\"jake\":\"stretchy\"}", json.GetFields())
}

func TestJSONIsEmpty(t *testing.T) {
	json := format.JSON()

	// should be true if no fields have been added
	assert.True(t, json.IsEmpty())

	json.AddField("bubblegum", "candy")
	// should be false if one or more fields have been added
	assert.False(t, json.IsEmpty())
}
