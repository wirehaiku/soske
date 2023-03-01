// Package clui implements command-line user interface functions.
package clui

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// GetPath returns the default database file path from a map of operating system names
// to {envvar, subfile} slices.
func GetPath(opts map[string][]string) (string, error) {
	if pair, ok := opts[runtime.GOOS]; ok {
		if dir, ok := os.LookupEnv(pair[0]); ok {
			return filepath.Join(dir, pair[1]), nil
		}

		return "", fmt.Errorf("environment variable %s not set", pair[0])
	}

	return "", fmt.Errorf("unsupported operating system")
}

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
