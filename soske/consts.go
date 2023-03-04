package soske

import "fmt"

const (
	// VersionDate is Soske's current version date.
	VersionDate = "2023-03-01"

	// VersionNums is Soske's current version number.
	VersionNums = "0.1.0"
)

var (
	// Debug indicates if Soske is in debug mode.
	Debug = false

	// Version is a description of Soske's current version.
	Version = fmt.Sprintf("Soske version %s (%s).", VersionNums, VersionDate)
)
