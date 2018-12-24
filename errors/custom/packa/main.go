package packa

import "fmt"

type temporaryError string

func (*temporaryError) Error() string {
	return "packa: temporary error"
}

func (*temporaryError) temporary() bool {
	return true
}

type basicError struct {
	value int
}

func (e *basicError) Error() string {
	return fmt.Sprintf("packa: value %v", e.value)
}
