package clean

import (
	"path/filepath"
	"strings"
)

// Key returns a clean lowercase key string.
func Key(key string) string {
	return strings.ToLower(strings.TrimSpace(key))
}

// Keys returns a slice of clean lowercase key strings.
func Keys(keys []string) []string {
	for n, key := range keys {
		keys[n] = Key(key)
	}

	return keys
}

// Path returns a clean file path.
func Path(path string) string {
	return filepath.Clean(path)
}

// Paths returns a slice of clean file paths.
func Paths(paths []string) []string {
	for n, path := range paths {
		paths[n] = Path(path)
	}

	return paths
}

// Value returns a clean value string.
func Value(val string) string {
	return strings.TrimSpace(val)
}

// Values returns a slice of clean value strings.
func Values(vals []string) []string {
	for n, val := range vals {
		vals[n] = Value(val)
	}

	return vals
}
