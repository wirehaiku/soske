package soske

import (
	"path/filepath"
	"strings"
)

// CleanKey returns a clean lowercase key string.
func CleanKey(key string) string {
	return strings.ToLower(strings.TrimSpace(key))
}

// CleanKeys returns a slice of clean lowercase key strings.
func CleanKeys(keys []string) []string {
	for n, key := range keys {
		keys[n] = CleanKey(key)
	}

	return keys
}

// CleanPath returns a clean file path.
func CleanPath(path string) string {
	return filepath.Clean(path)
}

// CleanPaths returns a slice of clean file paths.
func CleanPaths(paths []string) []string {
	for n, path := range paths {
		paths[n] = CleanPath(path)
	}

	return paths
}

// CleanValue returns a clean value string.
func CleanValue(value string) string {
	return strings.TrimSpace(value)
}

// CleanValues returns a slice of clean value strings.
func CleanValues(values []string) []string {
	for n, value := range values {
		values[n] = CleanValue(value)
	}

	return values
}
