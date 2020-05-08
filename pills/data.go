package main

import (
	"fmt"
	"sort"
	"strings"
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
	s := make([]string, len(d))

	for i := range d {
		s[i] = d[i].String()
	}

	return strings.Join(s, "")
}

func sortData(d Data) Data {
	sort.Slice(d, func(i, j int) bool {
		return d[i].DaysToExpire < d[j].DaysToExpire
	})

	return d
}
