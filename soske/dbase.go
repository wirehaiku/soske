package soske

import (
	"fmt"
	"sort"

	"go.etcd.io/bbolt"
)

// Connect returns a connected Bolt database.
func Connect(path string) (*bbolt.DB, error) {
	db, err := bbolt.Open(path, 0755, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot open database %q", path)
	}

	return db, nil
}

// DelKey deletes an existing key from a bucket.
func DelKey(db *bbolt.DB, bucket, key string) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		obj, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		return obj.Delete([]byte(key))
	})

	if err != nil {
		return fmt.Errorf("cannot update database %q", db.Path())
	}

	return nil
}

// GetKey returns the value of an existing key in a bucket.
func GetKey(db *bbolt.DB, bucket, key string) (string, error) {
	var value string
	err := db.View(func(tx *bbolt.Tx) error {
		obj := tx.Bucket([]byte(bucket))
		value = string(obj.Get([]byte(key)))
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("cannot read database %q", db.Path())
	}

	return CleanValue(value), nil
}

// ListKeys returns all existing keys in a bucket.
func ListKeys(db *bbolt.DB, bucket string) ([]string, error) {
	var keys []string
	err := db.View(func(tx *bbolt.Tx) error {
		obj := tx.Bucket([]byte(bucket))
		return obj.ForEach(func(key, _ []byte) error {
			keys = append(keys, string(key))
			return nil
		})
	})

	if err != nil {
		return nil, fmt.Errorf("cannot read database %q", db.Path())
	}

	sort.Strings(keys)
	return keys, nil
}

// SetKey sets the value of a new or existing key in a bucket.
func SetKey(db *bbolt.DB, bucket, key, value string) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		obj, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		return obj.Put([]byte(key), []byte(CleanValue(value)))
	})

	if err != nil {
		return fmt.Errorf("cannot Update database %q", db.Path())
	}

	return nil
}
