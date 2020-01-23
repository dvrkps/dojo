// Package authorization provides api authorization logic.
package authorization

import (
	"errors"
	"fmt"
	"strings"
)

const headerKey = "Authorization"

type header struct {
	algorithm algorithm
	keys      []string
	signature string
	apiKey    string
}

func (h *header) testString() string {
	return fmt.Sprintf("algorithm:%q keys:%q signature:%q apiKey:%q",
		h.algorithm,
		h.keys,
		h.signature,
		h.apiKey,
	)
}

func newHeader(value string) (*header, error) {
	if value == "" {
		return nil, errors.New("nil or empty value")
	}

	unquoted := strings.ReplaceAll(value, `"`, "")

	fields := strings.Split(unquoted, ",")

	var h header

	for i := range fields {
		kv := strings.SplitN(fields[i], "=", 2)
		switch kv[0] {
		case "algorithm":
			h.algorithm = algorithm(kv[1])
		case "headers":
			h.keys = strings.Fields(kv[1])
		case "signature":
			h.signature = kv[1]
		case "apikey":
			h.apiKey = kv[1]
		}
	}

	return &h, nil
}
