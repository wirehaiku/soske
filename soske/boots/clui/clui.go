// Package clui implements command-line user interface functions.
package clui

import (
	"strings"
)

// ParseArgs returns a parsed command name and argument slice from an argument slice.
func ParseArgs(args []string) (string, []string) {
	switch len(args) {
	case 0:
		return "", nil
	case 1:
		return strings.TrimSpace(args[0]), nil
	default:
		return strings.TrimSpace(args[0]), args[1:]
	}
}
