package main

import "testing"

type newHeightTestCase struct {
	ok bool
	in float64
}

func newHeightTestCases() map[string]newHeightTestCase {
	const cm = 0.01
	c := map[string]newHeightTestCase{
		"ok":            {ok: true, in: minHeight + cm},
		"min":           {ok: true, in: minHeight},
		"max":           {ok: true, in: maxHeight},
		"zero":          {in: 0},
		"less than min": {in: minHeight - cm},
		"more than max": {in: maxHeight + cm},
	}
	return c
}

func TestNewHeight(t *testing.T) {
	for name, tc := range newHeightTestCases() {
		t.Run(name, func(t *testing.T) {
			testNewHeight(t, tc)
		})
	}
}

func testNewHeight(t *testing.T, tc newHeightTestCase) {
	got, err := NewHeight(tc.in)
	if !tc.ok {
		want := Height(0)
		if got != want || err == nil {
			t.Errorf("NewHeight(%v) = %v, %v; want %v, <error>", tc.in, got, err, want)
		}
		return
	}
	want := Height(tc.in)
	if got != want || err != nil {
		t.Errorf("NewHeight(%v) = %v, %v; want %v, <nil>", tc.in, got, err, want)
	}
}
