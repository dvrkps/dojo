package packa

import "fmt"

type temporaryError string

// Error implements error interface.
func (*temporaryError) Error() string {
	return "packa: temporary error"
}

func (*temporaryError) temporary() bool {
	return true
}

type basicError struct {
	value int
}

// Error implements error interface.
func (e *basicError) Error() string {
	return fmt.Sprintf("packa: value %v", e.value)
}

// IsTemporary returns true if err is temporary.
func IsTemporary(err error) bool {
	type temporary interface {
		temporary() bool
	}
	te, ok := err.(temporary)
	return ok && te.Temporary()
}
