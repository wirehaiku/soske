// Package data implements data sanitisation and generation functions.
package data

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

// Body returns a clean plaintext body string.
func Body(body string) string {
	return strings.TrimSpace(body)
}

// Hash returns a SHA256 hash from a plaintext body string.
func Hash(body string) string {
	sum := sha256.Sum256([]byte(body))
	return fmt.Sprintf("%x", sum)
}

// Unix returns the current Unix timestamp as a string.
func Unix() string {
	unix := time.Now().Unix()
	return fmt.Sprintf("%d", unix)
}
