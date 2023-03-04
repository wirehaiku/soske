package soske

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestPackage(t *testing.T) {
	// success
	db := sqlx.MustConnect("sqlite3", ":memory:")
	_, err := db.Exec(Pragma + Schema)
	assert.NoError(t, err)
}
