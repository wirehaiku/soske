package soske

import (
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestDbOpen(t *testing.T) {
	// setup
	file, _ := os.CreateTemp("", "soske-test-*.db")
	defer file.Close()
	defer os.Remove(file.Name())

	// success - pragma & schema
	var num int
	db := DbOpen(file.Name(), Pragma, "create table Test(foo)")
	db.Get(&num, "select count(*) from SQLITE_SCHEMA")
	assert.Equal(t, 1, num)
	assert.NotNil(t, db)

	// success - pragma only
	db = DbOpen(file.Name(), Pragma, "drop table Test")
	db.Get(&num, "select count(*) from SQLITE_SCHEMA")
	assert.Equal(t, 1, num)
	assert.NotNil(t, db)
}

func TestDbExecute(t *testing.T) {
	// setup
	db := TestDB()

	// success
	var str string
	DbExecute(db, "insert into Keys (name) values (?)", "test")
	db.Get(&str, "select name from Keys order by init desc", "test")
	assert.Equal(t, "test", str)
}

func TestDbInteger(t *testing.T) {
	// setup
	db := TestDB()

	// success
	num := DbInteger(db, "select init from Keys where name=?", "alpha")
	assert.Equal(t, 1000, num)
}

func TestDbString(t *testing.T) {
	// setuo
	db := TestDB()

	// success
	str := DbString(db, "select name from Keys where name=?", "alpha")
	assert.Equal(t, "alpha", str)
}

func TestDbStrings(t *testing.T) {
	// setuo
	db := TestDB()

	// success
	strs := DbStrings(db, "select name from Keys where name=?", "alpha")
	assert.Equal(t, []string{"alpha"}, strs)
}
