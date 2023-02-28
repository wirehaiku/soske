// Package test implements unit-testing definitions and helper functions.
package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

// TestData is a map of default testing data.
var TestData = map[string]map[string][]byte{
	"alpha": {
		"body": []byte("Alpha value."),
		"hash": []byte("49e8c3bb0a4c0773b54af4aee638ef128c5dceae19b2e5adba57f0bdc33d4840"),
		"time": []byte("943920000"),
	},
	"bravo": {
		"body": []byte("Bravo value."),
		"hash": []byte("e628d55d2c5c5e47bda1fbb4fe8c8a365eb12c89d2745346216e20cad0b4a0c3"),
		"time": []byte("943923600"),
	},
}

// AssertBucket asserts the contents of a database bucket.
func AssertBucket(t *testing.T, db *bbolt.DB, name string, bmap map[string][]byte) {
	db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte(name))
		assert.NotNil(t, buck)
		return buck.ForEach(func(key, val []byte) error {
			assert.Equal(t, bmap[string(key)], val)
			return nil
		})
	})
}

// AssertNoBucket asserts the non-existence of a database bucket.
func AssertNoBucket(t *testing.T, db *bbolt.DB, name string) {
	db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte(name))
		assert.Nil(t, buck)
		return nil
	})
}

// KillDB closes and deletes an open database.
func KillDB(db *bbolt.DB) {
	db.Close()
	os.Remove(db.Path())
}

// MakeDB returns an open database populated with test data.
func MakeDB() *bbolt.DB {
	file, _ := os.CreateTemp("", "soske-test-*.db")
	file.Close()

	db, _ := bbolt.Open(file.Name(), 0755, nil)
	db.Update(func(tx *bbolt.Tx) error {
		for name, bmap := range TestData {
			buck, _ := tx.CreateBucket([]byte(name))
			for key, val := range bmap {
				buck.Put([]byte(key), val)
			}
		}
		return nil
	})

	return db
}
