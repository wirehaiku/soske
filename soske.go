package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/wirehaiku/soske/soske"
)

func main() {
	var cmd string
	var args []string

	if len(os.Args) > 1 {
		cmd = strings.ToLower(os.Args[1])
	}

	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	if cfun, ok := soske.Cmds[cmd]; ok {
		soske.Debug = true
		path := os.ExpandEnv("$HOME/.soske")
		db := soske.DbOpen(path, soske.Pragma, soske.Schema)
		for _, out := range cfun(db, args) {
			fmt.Println(out)
		}
	} else {
		soske.Die("invalid command")
	}
}
