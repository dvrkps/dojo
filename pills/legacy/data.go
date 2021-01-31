package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/dvrkps/dojo/pills/legacy/medicament"
)

// Data holds pills data.
type Data []medicament.Medicament

// Add adds pill.
func (d Data) Add(in []byte, cd time.Time) (Data, error) {
	m, err := medicament.New(cd, in)
	if err != nil {
		return Data{}, fmt.Errorf("data: %v", err)
	}

	d = append(d, m)

	return d, nil
}

// String returns all pills.
func (d Data) String() string {
	s := make([]string, len(d))

	for i := range d {
		s[i] = d[i].String()
	}

	const newline = "\n"

	return strings.Join(s, newline) + newline
}

func sortData(d Data) Data {
	sort.Slice(d, func(i, j int) bool {
		return d[i].DaysToExpire < d[j].DaysToExpire
	})

	return d
}
