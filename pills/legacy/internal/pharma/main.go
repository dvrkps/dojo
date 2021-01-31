// Package pharma implements pills actions.
package pharma

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/dvrkps/dojo/pills/legacy/medicament"
)

// All parses pills from file.
func All(r io.Reader, t time.Time) ([]medicament.Medicament, error) {
	s := bufio.NewScanner(r)

	const commentPrefix = "//"

	var all []medicament.Medicament

	for s.Scan() {
		b := s.Bytes()

		if bytes.HasPrefix(b, []byte(commentPrefix)) {
			continue
		}

		m, err := medicament.New(t, b)
		if err != nil {
			return nil, err
		}

		all = append(all, m)
	}

	return all, s.Err()
}

// Display show all pills.
func Display(all []medicament.Medicament) string {
	s := make([]string, len(all))

	for i := range all {
		s[i] = all[i].String()
	}

	const newline = "\n"

	return strings.Join(s, newline) + newline
}
