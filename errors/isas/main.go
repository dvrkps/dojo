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

	return "e: error"
}

func (e *myError) Unwrap() error {
	return e.err
}
