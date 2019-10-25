package main

import "testing"

type newPersonTestCase struct {
	fail bool
	h    float64
	w    float64
}

func newPersonTestCases() map[string]newPersonTestCase {
	const (
		one = 0.01
	)
	var (
		hok = minHeight + one
		wok = minWeight + one
	)
	c := map[string]newPersonTestCase{
		"ok":      {h: hok, w: wok},
		"h < min": {fail: true, h: minHeight - one, w: wok},
		"h > max": {fail: true, h: maxHeight + one, w: wok},
		"h min":   {h: minHeight, w: wok},
		"h max":   {h: maxHeight, w: wok},
		"w < min": {fail: true, h: hok, w: minWeight - one},
		"w > max": {fail: true, h: hok, w: maxWeight + one},
		"w min":   {fail: false, h: hok, w: minWeight},
		"w max":   {fail: false, h: hok, w: maxWeight},
	}
	return c
}

func TestNewPerson(t *testing.T) {
	for name, tc := range newPersonTestCases() {
		t.Run(name, func(t *testing.T) {
			testNewPerson(t, tc)
		})
	}
}

func testNewPerson(t *testing.T, tc newPersonTestCase) {
	_, err := NewPerson(tc.h, tc.w)
	if tc.fail {
		if err == nil {
			t.Errorf("NewPerson(%v, %v) = _, %v; want _, <error>", tc.h, tc.w, err)
		}
		return
	}
	if err != nil {
		t.Errorf("NewPerson(%v, %v) = _, %v; want _, <nil>", tc.h, tc.w, err)
	}
}

func TestPerson_Err(t *testing.T) {
	for name, tc := range newPersonTestCases() {
		t.Run(name, func(t *testing.T) {
			testPerson_Err(t, tc)
		})
	}
}

func testPerson_Err(t *testing.T, tc newPersonTestCase) {
	p := Person{h: tc.h, w: tc.w}
	err := p.Err()
	if tc.fail {
		if err == nil {
			t.Errorf("Person(%v, %v).Err() = %v; want <error>", tc.h, tc.w, err)
		}
		return
	}
	if err != nil {
		t.Errorf("Person(%v, %v).Err() = %v; <nil>", tc.h, tc.w, err)
	}
}
