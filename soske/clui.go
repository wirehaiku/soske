package soske

import (
	"fmt"
	"os"
)

// Die fatally prints a formatted string.
func Die(str string, args ...any) {
	fmt.Printf(fmt.Sprintf("Error: %s.\n", str), args...)
	os.Exit(1)
}

// Try fatally prints a formatted string on a non-nil error. If Debug is true, the error
// is also printed.
func Try(err error, str string, args ...any) {
	if err != nil {
		fmt.Printf(fmt.Sprintf("Error: %s.\n", str), args...)
		if Debug {
			fmt.Printf("Debug: %s.\n", err.Error())
		}
		os.Exit(1)
	}
}
