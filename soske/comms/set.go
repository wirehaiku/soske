package comms

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

// SetCommand sets the value of a new or existing name.
func SetCommand(db *sqlx.DB, args []string) ([]string, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("not enough arguments")
	}

	name := strings.TrimSpace(args[0])
	body := strings.Join(args[1:], "\n")
	code := "insert or ignore into Keys (name) values (?)"
	if _, err := db.Exec(code, name); err != nil {
		return nil, err
	}

	code = "insert into Vals (key, body) values (?, ?)"
	if _, err := db.Exec(code, name, body); err != nil {
		return nil, err
	}

	return nil, nil
}
