package soske

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

func TestTestDB(t *testing.T) {
	// success - database exists
	db := TestDB()
	db.View(func(tx *bbolt.Tx) error {
		obj := tx.Bucket([]byte("storage"))
		value := obj.Get([]byte("alpha"))
		assert.Equal(t, []byte("Alpha value."), value)
		return nil
	})

	// success - database auto-deletes
	time.Sleep(250 * time.Millisecond)
	assert.NoFileExists(t, db.Path())
}
