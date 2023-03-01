package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/soske/soske/tools/data"
	"github.com/wirehaiku/soske/soske/tools/test"
)

func TestSetCommand(t *testing.T) {
	// setup
	db := test.DB()

	// success - new key
	outs, err := SetCommand(db, []string{"test", "one", "two"})
	body, _ := data.GetString(db, "select body from Vals where key='test' order by init, oid desc")
	assert.Empty(t, outs)
	assert.NoError(t, err)
	assert.Equal(t, "one\ntwo", body)

	// success - existing key
	outs, err = SetCommand(db, []string{"test", "three"})
	body, _ = data.GetString(db, "select body from Vals where key='test' order by init, oid desc")
	assert.Empty(t, outs)
	assert.NoError(t, err)
	assert.Equal(t, "three", body)
}
