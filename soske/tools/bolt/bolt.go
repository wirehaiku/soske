// Package bolt implements Bolt database handling functions.
package bolt

import (
	"fmt"

	"go.etcd.io/bbolt"
)

// Connect returns a connected Bolt database.
func Connect(path string) (*bbolt.DB, error) {
	db, err := bbolt.Open(path, 0755, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database %q: %w", path, err)
	}

	return db, nil
}

// Delete deletes an existing bucket from a database.
func Delete(db *bbolt.DB, bkt string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		if tx.Bucket([]byte(bkt)) != nil {
			if err := tx.DeleteBucket([]byte(bkt)); err != nil {
				return fmt.Errorf("cannot delete bucket %q: %w", bkt, err)
			}
		}

		return nil
	})
}

// Get returns an existing bucket from a database.
func Get(db *bbolt.DB, bkt string) (map[string]string, error) {
	bmap := make(map[string]string)
	return bmap, db.View(func(tx *bbolt.Tx) error {
		if obj := tx.Bucket([]byte(bkt)); obj != nil {
			err := obj.ForEach(func(key, val []byte) error {
				bmap[string(key)] = string(val)
				return nil
			})

			if err != nil {
				return fmt.Errorf("cannot get bucket %q: %w", bkt, err)
			}
		}

		return nil
	})
}

// List returns the names of all buckets in the database.
func List(db *bbolt.DB) ([]string, error) {
	var bkts []string
	return bkts, db.View(func(tx *bbolt.Tx) error {
		err := tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
			bkts = append(bkts, string(name))
			return nil
		})

		if err != nil {
			return fmt.Errorf("cannot list buckets: %w", err)
		}

		return nil
	})
}

// Set sets a new or existing bucket in the database.
func Set(db *bbolt.DB, bkt string, bmap map[string]string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		obj, err := tx.CreateBucketIfNotExists([]byte(bkt))
		if err != nil {
			return fmt.Errorf("cannot create bucket %q: %w", bkt, err)
		}

		for key, val := range bmap {
			if err := obj.Put([]byte(key), []byte(val)); err != nil {
				return fmt.Errorf("cannot set key in bucket %q: %w", bkt, err)
			}
		}

		return nil
	})
}
