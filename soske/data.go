package soske

import "github.com/jmoiron/sqlx"

// DbOpen returns a new database connection. If pragma is not nil it is executed before
// returning, if schema is not nil it is executed if the database has no tables.
func DbOpen(path, pragma, schema string) *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", path)
	Try(err, "cannot connect to database")

	if pragma != "" {
		_, err = db.Exec(pragma)
		Try(err, "cannot write to database")
	}

	if schema != "" {
		var num int
		err := db.Get(&num, "select count(*) from SQLITE_SCHEMA")
		Try(err, "cannot read from database")

		if num == 0 {
			_, err := db.Exec(schema)
			Try(err, "cannot write to database")
		}
	}

	return db
}

// DbExecute executes and commits a database query.
func DbExecute(db *sqlx.DB, sql string, args ...any) {
	_, err := db.Exec(sql, args...)
	Try(err, "cannot write to database")
}

// DbInteger returns the first resulting integer from a database query.
func DbInteger(db *sqlx.DB, sql string, args ...any) int {
	var num int
	err := db.Get(&num, sql, args...)
	Try(err, "cannot read from database")
	return num
}

// DbString returns the first resulting string from a database query.
func DbString(db *sqlx.DB, sql string, args ...any) string {
	var str string
	err := db.Get(&str, sql, args...)
	Try(err, "cannot read from database")
	return str
}

// DbStrings returns the all resulting strings from a database query.
func DbStrings(db *sqlx.DB, sql string, args ...any) []string {
	var strs []string
	err := db.Select(&strs, sql, args...)
	Try(err, "cannot read from database")
	return strs
}
