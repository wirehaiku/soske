package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/Soske/soske/tools/test"
)

func TestGetCommand(t *testing.T) {
	// setup
	db := test.DB()

	// success - populated keys
	outs, err := GetCommand(db, []string{"alpha", "bravo"})
	assert.Equal(t, []string{
		"Alpha two.\n",
		"Bravo one.\n",
	}, outs)
	assert.NoError(t, err)

	// success - empty key
	outs, err = GetCommand(db, []string{"charlie"})
	assert.Empty(t, outs)
	assert.NoError(t, err)
}
