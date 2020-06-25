package isas

import (
	"errors"
	"fmt"
)

var sentinelError = errors.New("sentinel error")

type myError struct {
	err error
}

func (e *myError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("my: %v", e.err)
	}

	return "my error"
}

func (e *myError) Unwrap() error {
	return e.err
}

type subError struct {
	err error
}

func (e *subError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("sub: %v", e.err)
	}

	return "sub error"
}

func (e *subError) Unwrap() error {
	return e.err
}
