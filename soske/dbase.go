package soske

import (
	"fmt"

	"go.etcd.io/bbolt"
)

// Connect returns a connected Bolt database.
func Connect(path string) (*bbolt.DB, error) {
	db, err := bbolt.Open(path, 0755, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database %q", path)
	}

	return db, nil
}

// DeleteBucket deletes an existing bucket from a database.
func DeleteBucket(db *bbolt.DB, bkt string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		if tx.Bucket([]byte(bkt)) != nil {
			if err := tx.DeleteBucket([]byte(bkt)); err != nil {
				return fmt.Errorf("cannot delete bucket %q", bkt)
			}
		}

		return nil
	})
}

// GetBucket returns an existing bucket from a database.
func GetBucket(db *bbolt.DB, bkt string) (map[string]string, error) {
	bmap := make(map[string]string)
	return bmap, db.View(func(tx *bbolt.Tx) error {
		if obj := tx.Bucket([]byte(bkt)); obj != nil {
			err := obj.ForEach(func(key, val []byte) error {
				bmap[string(key)] = string(val)
				return nil
			})

			if err != nil {
				return fmt.Errorf("cannot get bucket %q", bkt)
			}
		}

		return nil
	})
}

// ListBuckets returns the names of all buckets in the database.
func ListBuckets(db *bbolt.DB) ([]string, error) {
	var bkts []string
	return bkts, db.View(func(tx *bbolt.Tx) error {
		err := tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
			bkts = append(bkts, string(name))
			return nil
		})

		if err != nil {
			return fmt.Errorf("cannot list buckets")
		}

		return nil
	})
}

// SetBucket sets a new or existing bucket in the database.
func SetBucket(db *bbolt.DB, bkt string, bmap map[string]string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		obj, err := tx.CreateBucketIfNotExists([]byte(bkt))
		if err != nil {
			return fmt.Errorf("cannot create bucket %q", bkt)
		}

		for key, val := range bmap {
			if err := obj.Put([]byte(key), []byte(val)); err != nil {
				return fmt.Errorf("cannot set key in bucket %q", bkt)
			}
		}

		return nil
	})
}
