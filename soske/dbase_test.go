package soske

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

func TestConnect(t *testing.T) {
	// setup
	file, _ := os.CreateTemp("", "soske-test-db-*")
	defer file.Close()
	defer os.Remove(file.Name())

	// success
	db, err := Connect(file.Name())
	assert.NotNil(t, db)
	assert.NoError(t, err)
}

func TestDelKey(t *testing.T) {
	// setup
	db := TestDB()

	// success
	err := DelKey(db, "storage", "alpha")
	assert.NoError(t, err)
	db.View(func(tx *bbolt.Tx) error {
		obj := tx.Bucket([]byte("storage"))
		value := obj.Get([]byte("alpha"))
		assert.Empty(t, value)
		return nil
	})
}

func TestGetKey(t *testing.T) {
	// setup
	db := TestDB()

	// success
	value, err := GetKey(db, "storage", "alpha")
	assert.Equal(t, "Alpha value.", value)
	assert.NoError(t, err)
}

func TestListKeys(t *testing.T) {
	// setup
	db := TestDB()

	// success
	keys, err := ListKeys(db, "storage")
	assert.Equal(t, []string{"alpha", "bravo", "charlie"}, keys)
	assert.NoError(t, err)
}

func TestSetKey(t *testing.T) {
	// setup
	db := TestDB()

	// success
	err := SetKey(db, "storage", "alpha", "newvalue")
	assert.NoError(t, err)
	db.View(func(tx *bbolt.Tx) error {
		obj := tx.Bucket([]byte("storage"))
		value := obj.Get([]byte("alpha"))
		assert.Equal(t, []byte("newvalue"), value)
		return nil
	})
}
