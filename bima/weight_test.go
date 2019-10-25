package main

import "testing"

type newWeightTestCase struct {
	ok bool
	in float64
}

func newWeightTestCases() map[string]newWeightTestCase {
	const dag = 0.01
	c := map[string]newWeightTestCase{
		"ok":            {ok: true, in: minWeight + dag},
		"min":           {ok: true, in: minWeight},
		"max":           {ok: true, in: maxWeight},
		"zero":          {in: 0},
		"less than min": {in: minWeight - dag},
		"more than max": {in: maxWeight + dag},
	}
	return c
}

func TestNewWeight(t *testing.T) {
	for name, tc := range newWeightTestCases() {
		t.Run(name, func(t *testing.T) {
			testNewWeight(t, tc)
		})
	}
}

func testNewWeight(t *testing.T, tc newWeightTestCase) {
	got, err := NewWeight(tc.in)
	if !tc.ok {
		want := Weight(0)
		if got != want || err == nil {
			t.Errorf("NewWeight(%v) = %v, %v; want %v, <error>", tc.in, got, err, want)
		}
		return
	}
	want := Weight(tc.in)
	if got != want || err != nil {
		t.Errorf("NewWeight(%v) = %v, %v; want %v, <nil>", tc.in, got, err, want)
	}
}
