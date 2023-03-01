// Package comms implements high-level command functions.
package comms

import "github.com/jmoiron/sqlx"

// Command is a user-executable command function.
type Command func(*sqlx.DB, []string) ([]string, error)

// Commands is a map of all existing commands.
var Commands = map[string]Command{
	"get": GetCommand,
	"lst": LstCommand,
}

// Run runs a Command with arguments against a database.
func Run(db *sqlx.DB, name string, args []string) ([]string, error) {
	if comm, ok := Commands[name]; ok {
		return comm(db, args)
	}

	return nil, nil
}
