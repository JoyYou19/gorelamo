package errors

import (
	"errors"
)

var (
	// Generic fallback
	ErrUnknown = errors.New("contextprime: unknown error")
)

// DBError represents any error returned by the database
type DBError struct {
	Message string // decoded, human-readable
	Raw     []byte // original response body
}

func (e *DBError) Error() string {
	return e.Message
}
