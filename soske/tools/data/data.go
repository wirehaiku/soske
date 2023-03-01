// Package data implements low-level database functions.
package data

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Open returns an open database connection, executes pragma and initialises the
// database if it does not contain any schema.
func Open(path, pragma, schema string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(pragma); err != nil {
		return nil, err
	}

	num, err := GetInt(db, "select count(*) from SQLITE_SCHEMA")
	if err != nil {
		return nil, err
	}

	if num == 0 {
		if _, err = db.Exec(schema); err != nil {
			return nil, err
		}
	}

	return db, nil
}

// GetInt returns an integer from the first row of a database query.
func GetInt(db *sqlx.DB, code string, args ...any) (int, error) {
	var num int
	if err := db.Get(&num, code, args...); err != nil {
		return 0, err
	}

	return num, nil
}

// GetTime returns a Time object from an integer in the first row of a database query.
func GetTime(db *sqlx.DB, code string, args ...any) (time.Time, error) {
	var num int64
	if err := db.Get(&num, code, args...); err != nil {
		return time.Unix(0, 0), err
	}

	return time.Unix(num, 0), nil
}

// GetString returns a string from the first row of a database query.
func GetString(db *sqlx.DB, code string, args ...any) (string, error) {
	var str string
	if err := db.Get(&str, code, args...); err != nil {
		return "", err
	}

	return str, nil
}
