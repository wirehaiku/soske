package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

func TestAssertDB(t *testing.T) {
	// setup
	db := MakeDB()
	defer KillDB(db)

	// success
	AssertDB(t, db, "alpha", TestData["alpha"])
}

func TestKillDB(t *testing.T) {
	// setup
	db := MakeDB()

	// success
	KillDB(db)
	assert.NoFileExists(t, db.Path())
}

func TestMakeDB(t *testing.T) {
	// success
	db := MakeDB()
	db.View(func(tx *bbolt.Tx) error {
		obj := tx.Bucket([]byte("alpha"))
		val := obj.Get([]byte("body"))
		assert.Equal(t, []byte("Alpha value."), val)
		return nil
	})
}
