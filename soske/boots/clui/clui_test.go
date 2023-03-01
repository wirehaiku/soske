package clui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgs(t *testing.T) {
	// success - zero arguments
	name, args := ParseArgs([]string{})
	assert.Empty(t, name)
	assert.Empty(t, args)

	// success - one arguments
	name, args = ParseArgs([]string{"test"})
	assert.Equal(t, "test", name)
	assert.Empty(t, args)

	// success - multiple arguments
	name, args = ParseArgs([]string{"test", "one", "two"})
	assert.Equal(t, "test", name)
	assert.Equal(t, []string{"one", "two"}, args)
}
