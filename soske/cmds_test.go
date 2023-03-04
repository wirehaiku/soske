package soske

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCmd(t *testing.T) {
	// setup
	db := TestDB()

	// success
	outs := GetCmd(db, []string{"alpha", "bravo"})
	assert.Equal(t, []string{"Alpha two.", "Bravo one."}, outs)
}

func TestLstCmd(t *testing.T) {
	// setup
	db := TestDB()

	// success
	outs := LstCmd(db, []string{"a", "!charlie"})
	assert.Equal(t, []string{"alpha", "bravo"}, outs)
}

func TestSetCmd(t *testing.T) {
	// setup
	db := TestDB()

	// success
	outs := SetCmd(db, []string{"test", "one", "two"})
	body := DbString(db, "select body from GoodVals where key=?", "test")
	assert.Empty(t, outs)
	assert.Equal(t, "one\ntwo", body)
}
