package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/soske/soske/tools/test"
)

func TestLstCommand(t *testing.T) {
	// setup
	db := test.DB()

	// success - no arguments
	outs, err := LstCommand(db, nil)
	assert.Equal(t, []string{"alpha", "bravo"}, outs)
	assert.NoError(t, err)

	// success - positive argument
	outs, err = LstCommand(db, []string{"alp"})
	assert.Equal(t, []string{"alpha"}, outs)
	assert.NoError(t, err)

	// success - negative argument
	outs, err = LstCommand(db, []string{"!alp"})
	assert.Equal(t, []string{"bravo"}, outs)
	assert.NoError(t, err)
}
