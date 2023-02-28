package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tvalue\n")
	assert.Equal(t, "value", body)
}

func TestHash(t *testing.T) {
	// success
	hash := Hash("value")
	assert.Equal(t, "cd42404d52ad55ccfa9aca4adc828aa5800ad9d385a0671fbcbf724118320619", hash)
}

func TestUnix(t *testing.T) {
	// success
	unix := Unix()
	assert.Regexp(t, `\d+`, unix)
}
