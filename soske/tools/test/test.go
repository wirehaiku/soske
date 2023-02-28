// Package test implements unit-testing definitions and helper functions.
package test

import (
	"os"
	"time"

	"go.etcd.io/bbolt"
)

// Data is a map of default testing data.
var Data = map[string]map[string]string{
	"alpha": {
		"body": "Alpha value.",
		"hash": "49e8c3bb0a4c0773b54af4aee638ef128c5dceae19b2e5adba57f0bdc33d4840",
		"time": "943920000",
	},
	"bravo": {
		"body": "Bravo value.",
		"hash": "e628d55d2c5c5e47bda1fbb4fe8c8a365eb12c89d2745346216e20cad0b4a0c3",
		"time": "943923600",
	},
}

// DB returns an open Bolt database populated with test data. The database file is
// automatically deleted after 200 milliseconds.
func DB() *bbolt.DB {
	file, _ := os.CreateTemp("", "soske-test-*.db")
	defer file.Close()

	db, _ := bbolt.Open(file.Name(), 0755, nil)
	db.Update(func(tx *bbolt.Tx) error {
		for bkt, bmap := range Data {
			obj, _ := tx.CreateBucket([]byte(bkt))
			for key, val := range bmap {
				obj.Put([]byte(key), []byte(val))
			}
		}
		return nil
	})

	go func() {
		time.Sleep(200 * time.Millisecond)
		os.Remove(file.Name())
	}()

	return db
}
