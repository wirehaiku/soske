package sqls

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestPackage(t *testing.T) {
	// success - production schema
	db := sqlx.MustConnect("sqlite3", ":memory:")
	_, err := db.Exec(BasePragma + ProdSchema)
	assert.NoError(t, err)

	// success - testing schema
	db = sqlx.MustConnect("sqlite3", ":memory:")
	_, err = db.Exec(BasePragma + TestSchema)
	assert.NoError(t, err)
}
