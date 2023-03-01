// Soske: Stephen's Old-School Key Engine.
package main

import (
	"fmt"
	"os"

	"github.com/wirehaiku/soske/soske/comms"
	"github.com/wirehaiku/soske/soske/tools/clui"
	"github.com/wirehaiku/soske/soske/tools/data"
	"github.com/wirehaiku/soske/soske/tools/defs"
	"github.com/wirehaiku/soske/soske/tools/sqls"
)

// die fatally prints a formatted string.
func die(str string, args ...any) {
	fmt.Printf(fmt.Sprintf("Error: %s.\n", str), args...)
	os.Exit(1)
}

// try fatally prints a non-nil error.
func try(err error) {
	if err != nil {
		die(err.Error())
	}
}

// main runs the main Soske program.
func main() {
	// Get command-line arguments.
	name, args := clui.ParseArgs(os.Args[1:])

	// Get default database path.
	path, err := clui.GetPath(defs.DefaultPaths)
	try(err)

	// Open database file.
	db, err := data.Open(path, sqls.BasePragma, sqls.ProdSchema)
	if err != nil {
		die("cannot access database (%s)", err.Error())
	}

	// Execute argued command.
	outs, err := comms.Run(db, name, args)
	if err != nil {
		die("%s command failed (%s)", name, err.Error())
	}

	// Print command results.
	for _, out := range outs {
		fmt.Println(out)
	}
}
