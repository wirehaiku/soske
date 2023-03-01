package data

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/soske/soske/tools/sqls"
	"github.com/wirehaiku/soske/soske/tools/test"
)

func TestOpen(t *testing.T) {
	// setup
	file, _ := os.CreateTemp("", "soske-test-*.db")
	defer file.Close()
	defer os.Remove(file.Name())

	// success - first time initialised
	var num int
	db, err := Open(file.Name(), sqls.BasePragma, sqls.TestSchema)
	db.Get(&num, "select count(*) from Keys")
	assert.Equal(t, 3, num)
	assert.NoError(t, err)

	// success - second time not initialised
	db, err = Open(file.Name(), sqls.BasePragma, sqls.TestSchema)
	db.Get(&num, "select count(*) from Keys")
	assert.Equal(t, 3, num)
	assert.NoError(t, err)
}

func TestGetInt(t *testing.T) {
	// setup
	db := test.DB()

	// success
	num, err := GetInt(db, "select init from Keys where name=?", "alpha")
	assert.Equal(t, 100, num)
	assert.NoError(t, err)
}

func TestGetTime(t *testing.T) {
	// setup
	db := test.DB()

	// success
	tme, err := GetTime(db, "select init from Keys where name=?", "alpha")
	assert.Equal(t, time.Unix(100, 0), tme)
	assert.NoError(t, err)
}

func TestGetString(t *testing.T) {
	// setup
	db := test.DB()

	// success
	str, err := GetString(db, "select name from Keys where name=?", "alpha")
	assert.Equal(t, "alpha", str)
	assert.NoError(t, err)
}
