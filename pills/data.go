package main

import (
	"bytes"
	"fmt"
	"sort"
	"time"

	"github.com/dvrkps/dojo/pills/medicament"
)

// Data holds pills data.
type Data []medicament.Medicament

// Add adds pill.
func (d *Data) Add(in []byte, cd time.Time) error {
	m, err := medicament.New(cd, in)
	if err != nil {
		return fmt.Errorf("data: %v", err)
	}

	*d = append(*d, m)

	return nil
}

// String return formated list of pills.
func (d Data) String() string {
	buf := bytes.NewBuffer(nil)

	for _, p := range d {
		_, _ = fmt.Fprint(buf, p)
	}

	return buf.String()
}

func sortData(d Data) Data {
	sort.Slice(d, func(i, j int) bool {
		return d[i].DaysToExpire < d[j].DaysToExpire
	})

	return d
}
