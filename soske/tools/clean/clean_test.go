package clean

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKey(t *testing.T) {
	// success
	key := Key("\tKEY\n")
	assert.Equal(t, "key", key)
}

func TestKeys(t *testing.T) {
	// success
	keys := Keys([]string{"\tKEY\n"})
	assert.Equal(t, []string{"key"}, keys)
}

func TestPath(t *testing.T) {
	// success
	path := Path("/././path.test")
	assert.Equal(t, "/path.test", path)
}

func TestPaths(t *testing.T) {
	// success
	paths := Paths([]string{"/././path.test"})
	assert.Equal(t, []string{"/path.test"}, paths)
}

func TestValue(t *testing.T) {
	// success
	val := Value("\tvalue\n")
	assert.Equal(t, "value", val)
}

func TestValues(t *testing.T) {
	// success
	vals := Values([]string{"\tvalue\n"})
	assert.Equal(t, []string{"value"}, vals)
}
