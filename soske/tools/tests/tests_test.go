package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

func TestDB(t *testing.T) {
	// success - database exists
	db := DB()
	db.View(func(tx *bbolt.Tx) error {
		obj := tx.Bucket([]byte("alpha"))
		val := obj.Get([]byte("body"))
		assert.Equal(t, []byte("Alpha value."), val)
		return nil
	})

	// success - database auto-deletes
	time.Sleep(250 * time.Millisecond)
	assert.NoFileExists(t, db.Path())
}
