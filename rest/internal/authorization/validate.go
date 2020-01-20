package authorization

import (
	"errors"
	"fmt"
	"net/http"
)

const headerKey = "Authorization"

func Validate(header http.Header) error {
	if len(header) == 0 {
		return errors.New("nil or empty header")
	}

	f, err := parseAuthField(header.Get(headerKey))
	if err != nil {
		return fmt.Errorf("field: %v", err)
	}

	_ = f
	return nil
}
