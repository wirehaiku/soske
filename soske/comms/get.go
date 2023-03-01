package comms

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/wirehaiku/soske/soske/tools/data"
)

// GetCommand returns the value of each existing key.
func GetCommand(db *sqlx.DB, args []string) ([]string, error) {
	var outs []string

	for _, key := range args {
		code := "select body from Vals where key=? order by init desc limit 1"
		body, err := data.GetString(db, code, key)
		if err != nil && err != sql.ErrNoRows {
			return nil, fmt.Errorf("cannot access database")
		}

		if body != "" {
			outs = append(outs, strings.TrimSpace(body)+"\n")
		}
	}

	return outs, nil
}
