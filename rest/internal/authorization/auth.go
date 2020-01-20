// Package authorization provides api authorization logic.
package authorization

import (
	"errors"
	"fmt"
	"strings"
)

type authField struct {
	algorithm  string
	headerKeys []string
	signature  string
	apiKey     string
}

func (af *authField) testString() string {
	return fmt.Sprintf("algorithm:%q headerKeys:%q signature:%q apiKey:%q",
		af.algorithm,
		af.headerKeys,
		af.signature,
		af.apiKey,
	)
}

func parseAuthField(value string) (*authField, error) {
	if value == "" {
		return nil, errors.New("nil or empty authorization field")
	}

	unquoted := strings.ReplaceAll(value, `"`, "")

	af := authField{}

	fields := strings.Split(unquoted, ",")
	for i := range fields {
		kv := strings.SplitN(fields[i], "=", 2)
		switch kv[0] {
		case "algorithm":
			af.algorithm = kv[1]
		case "headers":
			af.headerKeys = strings.Fields(kv[1])
		case "signature":
			af.signature = kv[1]
		case "apikey":
			af.apiKey = kv[1]
		}
	}

	switch af.algorithm {
	case "hmac-sha256":
	default:
		return nil, fmt.Errorf("invalid algorithm %q", af.algorithm)
	}

	if len(af.headerKeys) == 0 {
		return nil, errors.New("empty headers")
	}

	if af.signature == "" {
		return nil, errors.New("empty signature")
	}

	if af.apiKey == "" {
		return nil, errors.New("empty apikey")
	}

	return &af, nil
}
