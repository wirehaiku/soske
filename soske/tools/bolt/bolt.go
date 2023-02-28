// Package bolt implements Bolt database handler functions.
package bolt

import "go.etcd.io/bbolt"

// DeleteBucket deletes an existing bucket from a database.
func DeleteBucket(db *bbolt.DB, name string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(name))
	})
}

// GetBucket returns an existing bucket as a string/byteslice map.
func GetBucket(db *bbolt.DB, name string) (map[string][]byte, error) {
	bmap := make(map[string][]byte)
	return bmap, db.View(func(tx *bbolt.Tx) error {
		if buck := tx.Bucket([]byte(name)); buck != nil {
			return buck.ForEach(func(key, val []byte) error {
				bmap[string(key)] = val
				return nil
			})
		}
		return nil
	})
}

// GetKey returns a value from an existing bucket as a byteslice.
func GetKey(db *bbolt.DB, name, key string) ([]byte, error) {
	var val []byte
	return val, db.View(func(tx *bbolt.Tx) error {
		if buck := tx.Bucket([]byte(name)); buck != nil {
			val = buck.Get([]byte(key))
		}
		return nil
	})
}

// ListBuckets returns the names of all existing buckets in a database.
func ListBuckets(db *bbolt.DB) ([]string, error) {
	var names []string
	return names, db.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
			names = append(names, string(name))
			return nil
		})
	})
}

// SetBucket sets a new or existing bucket in a database from a string/byteslice map.
func SetBucket(db *bbolt.DB, name string, bmap map[string][]byte) error {
	return db.Update(func(tx *bbolt.Tx) error {
		buck, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}
		for key, val := range bmap {
			if err := buck.Put([]byte(key), val); err != nil {
				return err
			}
		}
		return nil
	})
}
