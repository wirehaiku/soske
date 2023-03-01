package comms

import (
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
)

// LstCommand returns the name of each existing key.
func LstCommand(db *sqlx.DB, args []string) ([]string, error) {
	var keys []string

	subq := "select count(*) from Vals where key=name"
	code := "select name from Keys where (" + subq + ")>0 order by name asc"
	if err := db.Select(&keys, code); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	for _, arg := range args {
		bok := func(b bool) bool { return !b }
		if strings.HasPrefix(arg, "!") {
			arg = strings.TrimLeft(arg, "!")
			bok = func(b bool) bool { return b }
		}

		for n := 0; n < len(keys); n++ {
			if bok(strings.Contains(keys[n], arg)) {
				keys = append(keys[:n], keys[1+n:]...)
				n--
			}
		}
	}

	return keys, nil
}
