package soske

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanKey(t *testing.T) {
	// success
	key := CleanKey("\tFOO\n")
	assert.Equal(t, "foo", key)
}

func TestCleanKeys(t *testing.T) {
	// success
	keys := CleanKeys([]string{"\tFOO\n"})
	assert.Equal(t, []string{"foo"}, keys)
}

func TestCleanPath(t *testing.T) {
	// success
	path := CleanPath("/././path.test")
	assert.Equal(t, "/path.test", path)
}

func TestCleanPaths(t *testing.T) {
	// success
	paths := CleanPaths([]string{"/././path.test"})
	assert.Equal(t, []string{"/path.test"}, paths)
}

func TestCleanValue(t *testing.T) {
	// success
	value := CleanValue("\tvalue\n")
	assert.Equal(t, "value", value)
}

func TestCleanValues(t *testing.T) {
	// success
	values := CleanValues([]string{"\tvalue\n"})
	assert.Equal(t, []string{"value"}, values)
}
