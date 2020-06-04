// Package pharma implements pills actions.
package pharma

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/dvrkps/dojo/pills/medicament"
)

// All holds pills.
type All struct {
	pills []medicament.Medicament
}

// NewAll parses pills from file.
func NewAll(r io.Reader, t time.Time) (*All, error) {
	var all All
	err := all.scan(r, t)
	if err != nil {
		return nil, err
	}
	return &all, nil
}

// String returns all pills string representation.
func (a *All) String() string {
	s := make([]string, len(a.pills))

	for i := range a.pills {
		s[i] = a.pills[i].String()
	}

	const newline = "\n"

	return strings.Join(s, newline) + newline
}

func (a *All) scan(r io.Reader, t time.Time) error {
	s := bufio.NewScanner(r)

	const commentPrefix = "//"

	for s.Scan() {
		if bytes.HasPrefix(s.Bytes(), []byte(commentPrefix)) {
			continue
		}

		if err := a.add(s.Bytes(), t); err != nil {
			return err
		}
	}

	return s.Err()
}

func (a *All) add(in []byte, t time.Time) error {
	m, err := medicament.New(t, in)
	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	a.pills = append(a.pills, m)

	return nil
}
