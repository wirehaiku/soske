package soske

import (
	"os"
	"time"

	"go.etcd.io/bbolt"
)

// TestData is a map of default testing data.
var TestData = map[string]map[string]string{
	"storage": {
		"alpha":   "Alpha value.",
		"bravo":   "Bravo value.",
		"charlie": "Charlie value.",
	},
}

// TestDB returns an open Bolt database populated with test data. The database file is
// automatically deleted after 200 milliseconds.
func TestDB() *bbolt.DB {
	path, _ := os.CreateTemp("", "soske-test-db-*")
	defer path.Close()

	db, _ := bbolt.Open(path.Name(), 0755, nil)
	db.Update(func(tx *bbolt.Tx) error {
		for bucket, keymap := range TestData {
			obj, _ := tx.CreateBucket([]byte(bucket))
			for key, value := range keymap {
				obj.Put([]byte(key), []byte(value))
			}
		}
		return nil
	})

	go func() {
		time.Sleep(200 * time.Millisecond)
		os.Remove(path.Name())
	}()

	return db
}
