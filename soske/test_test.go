package soske

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestDB(t *testing.T) {
	// success
	db := TestDB()
	assert.NotNil(t, db)
}
