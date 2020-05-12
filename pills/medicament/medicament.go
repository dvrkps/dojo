// Package medicament provides medicament functions.
package medicament

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

// Medicament holds medicament data.
type Medicament struct {
	refill
	Name string
	dosage
	expire
}

// String returns medicament string representation.
func (m Medicament) String() string {
	return fmt.Sprintf("%3d   %6s %3s   %s %v",
		m.DaysToExpire,
		m.ExpireDate.Format("2.1."),
		m.ExpireDate.Format("Mon"),
		m.Name,
		m.dosage)
}

// Compare compares medicaments by days to expiration.
func Compare(m1, m2 Medicament) bool {
	return m1.DaysToExpire < m2.DaysToExpire
}

// midnight returns date with zeroed time.
func midnight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

// New create medicament.
func New(date time.Time, in []byte) (Medicament, error) {
	empty := Medicament{}
	// format
	all := bytes.Split(in, []byte{','})

	const parts = 4
	if len(all) < parts {
		return empty, errors.New("medicament: format error")
	}
	// init
	m := Medicament{}
	// refill
	r, err := newRefill(all[0], all[1])
	if err != nil {
		return empty, fmt.Errorf("medicament: %s", err)
	}

	m.refill = r
	// name
	n, err := parseName(all[2])
	if err != nil {
		return empty, err
	}

	m.Name = n

	// dosage
	d, err := newDosage(all[3:]...)
	if err != nil {
		return empty, fmt.Errorf("medicament: %s", err)
	}

	m.dosage = d

	// expire
	e := newExpire(date, r, d.ratio())
	m.expire = e

	return m, nil
}

// parseName returns medicament name.
func parseName(in []byte) (string, error) {
	s := string(in)
	if s == "" {
		return "", errors.New("medicament: empty name")
	}

	return s, nil
}
