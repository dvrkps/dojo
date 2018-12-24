package packa

import "fmt"

// New creates packa.
func New(x int, max int) error {
	if x < 0 {
		return &temporaryError{value: x}
	}
	if x < max {
		return &basicError{value: x, max: max}
	}
	return nil
}

type temporaryError struct {
	value int
}

// Error implements error interface.
func (e *temporaryError) Error() string {
	return fmt.Sprintf("packa: temporary: value %v", e.value)
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
