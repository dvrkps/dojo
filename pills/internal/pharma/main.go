// Package pharma implements pills actions.
package pharma

import (
	"io"
	"time"

	"github.com/dvrkps/dojo/pills/medicament"
)

type Pharma struct {
	pills []medicament.Medicament
}

func New(t time.Time, r io.Reader) (*Pharma, error) {
	return nil, nil
}
