// Package defs implements program constants and definitions.
package defs

import "fmt"

const (
	// VersionDate is the date of Soske's current version.
	VersionDate = "2022-03-01"

	// VersionNums is the number of Soske's current version.
	VersionNums = "0.1.0"
)

var (
	// DefaultPaths is a map of operating systems to default storage paths.
	DefaultPaths = map[string][]string{
		"linux":   {"XDG_CONFIG_HOME", "soske/soske.db"},
		"windows": {"APPDATA", "soske/soske.db"},
	}

	// Version is a description of Soske's current version.
	Version = fmt.Sprintf("Soske version %s (%s).", VersionNums, VersionDate)
)
