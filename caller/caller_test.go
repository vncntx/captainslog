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
