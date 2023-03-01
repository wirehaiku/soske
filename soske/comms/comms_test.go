package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/soske/soske/tools/test"
)

func TestRun(t *testing.T) {
	// setup
	db := test.DB()

	// success
	outs, err := Run(db, "get", []string{"alpha"})
	assert.Equal(t, []string{"Alpha two.\n"}, outs)
	assert.NoError(t, err)
}
