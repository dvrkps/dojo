package authorization

import (
	"errors"
	"fmt"
	"net/http"
)

// Validate checks authorization header and signature.
func Validate(all http.Header) error {
	if len(all) == 0 {
		return errors.New("nil or empty http headers")
	}

	h, err := newHeader(all.Get(headerKey))
	if err != nil {
		return fmt.Errorf("header: %v", err)
	}

	s := signature{
		header:    all,
		algorithm: h.algorithm,
		keys:      h.keys,
		value:     h.signature,
	}

	return s.check()
}
