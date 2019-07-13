package caller_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vincentfiestada/captainslog/caller"
)

func TestGetName(t *testing.T) {
	assert := assert.New(t)

	expected := "github.com/vincentfiestada/captainslog/caller_test.TestGetName"
	assert.Equal(expected, caller.GetName(1))
	func() {
		assert.Equal(expected, caller.GetName(2))
	}()
}

func TestShorten(t *testing.T) {
	assert := assert.New(t)

	path := "github.com/vincentfiestada/captainslog/caller_test.TestShorten"
	assert.Equal("TestShorten", caller.Shorten(path, 15))
	assert.Equal("caller_test.TestShorten", caller.Shorten(path, 30))
}
