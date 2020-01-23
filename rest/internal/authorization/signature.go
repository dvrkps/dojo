package authorization

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type signature struct {
	header    http.Header
	algorithm algorithm
	keys      []string
	value     string
}

func (s *signature) check() error {
	if len(s.header) == 0 {
		return errors.New("nil or empty header")
	}

	msg, err := s.message()
	if err != nil {
		return fmt.Errorf("message: %v", err)
	}

	a := s.algorithm

	return a.check(msg, s.value)
}

func (s *signature) message() (string, error) {
	lines := make([]string, 0, len(s.keys))

	for _, k := range s.keys {
		v := s.header.Get(k)
		if v == "" {
			return "", fmt.Errorf("%q: empty key", k)
		}

		row := fmt.Sprintf("%v: %v", k, v)

		lines = append(lines, row)
	}

	const title = "title"

	lines = append([]string{title}, lines...)

	join := strings.Join(lines, "\n")

	return join, nil
}
