package clui

import (
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPath(t *testing.T) {
	// setup
	opts := map[string][]string{runtime.GOOS: {"TEST", "test.file"}}
	os.Setenv("TEST", "/test")

	// success
	path, err := GetPath(opts)
	assert.Equal(t, "/test/test.file", path)
	assert.NoError(t, err)
}

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
