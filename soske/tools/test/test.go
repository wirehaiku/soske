// Package test implements unit-testing helper functions.
package test

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wirehaiku/soske/soske/tools/sqls"
)

// DB returns an in-memory database populated with test data.
func DB() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", ":memory:")
	db.MustExec(sqls.BasePragma + sqls.TestSchema)
	return db
}
