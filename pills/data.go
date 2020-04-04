package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/dvrkps/dojo/pills/medicament"
)

// Data holds pills data.
type Data []*medicament.Medicament

// Add adds pill.
func (d *Data) Add(in []byte, cd time.Time) error {
	m, err := medicament.New(cd, in)
	if err != nil {
		return fmt.Errorf("data: %s", err)
	}

	*d = append(*d, &m)

	return nil
}

// Len returns Data length.
func (d Data) Len() int {
	return len(d)
}

// Less holds custom sort logic.
func (d Data) Less(i, j int) bool {
	return d[i].DaysToExpire < d[j].DaysToExpire
}

// String return formated list of pills.
func (d Data) String() string {
	buf := bytes.NewBuffer(nil)

	for _, p := range d {
		_, _ = fmt.Fprint(buf, p)
	}

	return buf.String()
}

// Swap swaps elements.
func (d Data) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
