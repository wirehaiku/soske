package soske

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Cmd is a user-executable command function.
type Cmd func(*sqlx.DB, []string) []string

// Cmds is a map of all existing commands.
var Cmds = map[string]Cmd{
	"get": GetCmd,
	"lst": LstCmd,
	"set": SetCmd,
}

// GetCmd prints the latest value of an existing key.
func GetCmd(db *sqlx.DB, args []string) []string {
	var outs []string
	for _, key := range args {
		body := DbString(db, "select body from GoodVals where key=?", key)
		outs = append(outs, body)
	}

	return outs
}

// LstCmd prints a list of existing keys.
func LstCmd(db *sqlx.DB, args []string) []string {
	if len(args) > 0 {
		var opts []string
		for _, sub := range args {
			if strings.HasPrefix(sub, "!") {
				sub = strings.TrimLeft(sub, "!")
				opts = append(opts, fmt.Sprintf("name not like '%%%s%%'", sub))
			} else {
				opts = append(opts, fmt.Sprintf("name like '%%%s%%'", sub))
			}
		}

		whr := strings.Join(opts, " and ")
		return DbStrings(db, "select name from Goodkeys where "+whr)
	}

	return DbStrings(db, "select name from GoodKeys")
}

// SetCmd sets the value of a new or existing key.
func SetCmd(db *sqlx.DB, args []string) []string {
	if len(args) < 2 {
		Die("not enough arguments")
	}

	key := strings.ToLower(args[0])
	val := strings.Join(args[1:], "\n")

	if val != DbString(db, "select body from GoodVals where key=?", key) {
		DbExecute(db, "insert or ignore into Keys (name) values (?)", key)
		DbExecute(db, "insert into Vals (key, body) values (?, ?)", key, val)
	}

	return nil
}
