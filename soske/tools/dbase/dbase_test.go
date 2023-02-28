package dbase

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/Soske/soske/tools/tests"
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

func TestDeleteBucket(t *testing.T) {
	// setup
	db := tests.DB()

	// success
	err := Delete(db, "alpha")
	assert.NoError(t, err)
	db.View(func(tx *bbolt.Tx) error {
		obj := tx.Bucket([]byte("alpha"))
		assert.Nil(t, obj)
		return nil
	})
}

func TestGetBucket(t *testing.T) {
	// setup
	db := tests.DB()

	// success
	bmap, err := Get(db, "alpha")
	assert.Equal(t, tests.Data["alpha"], bmap)
	assert.NoError(t, err)
}

func TestListBuckets(t *testing.T) {
	// setup
	db := tests.DB()

	// success
	bkts, err := List(db)
	assert.Equal(t, []string{"alpha", "bravo"}, bkts)
	assert.NoError(t, err)
}

func TestSetBucket(t *testing.T) {
	// setup
	db := tests.DB()

	// success
	err := Set(db, "test", map[string]string{"key": "value"})
	assert.NoError(t, err)
	db.View(func(tx *bbolt.Tx) error {
		obj := tx.Bucket([]byte("test"))
		val := obj.Get([]byte("key"))
		assert.Equal(t, []byte("value"), val)
		return nil
	})
}
