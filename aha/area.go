package main

import (
	"fmt"
	"math"
)

// Area in square meters.
type Area struct {
	value float64
	err   error
}

const (
	m2Kind  = "m2"
	ralKind = "ral"
	chvKind = "chv"
)

// NewArea create Area.
func NewArea(m2, ral, chv float64) (*Area, error) {
	a := Area{}
	a.add(m2Kind, m2)
	a.add(ralKind, ral)
	a.add(chvKind, chv)

	if a.err != nil {
		return nil, a.err
	}

	return &a, nil
}

// String returns formated area.
func (a *Area) String() string {
	m2 := a.value
	ral := math.Floor(m2 / oneRal)
	d := m2 - math.Ceil(ral*oneRal)

	var chv float64
	if d >= oneChv {
		chv = math.Ceil(d / oneChv)
	}

	const rowFmt = "%v m2 = %v ral, %v chv"

	return fmt.Sprintf(rowFmt, m2, ral, chv)
}

const (
	// oneRal is value of 1 Ral(kj) in m2.
	oneRal = 5754.542
	// oneChv is value of 1 square hvat(chv) in m2.
	oneChv = 3.596652
)

func (a *Area) add(kind string, value float64) {
	if a.err != nil {
		return
	}

	if value < 0 {
		a.err = fmt.Errorf("%v < 0", kind)
		return
	}

	switch kind {
	case m2Kind:
		a.value += value
		return
	case ralKind:
		value *= oneRal
	case chvKind:
		value *= oneChv
	}

	a.add("m2", value)
}
