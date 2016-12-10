package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []int{1, 2, 3}
	testAdd(t, value, tests)
}

func testAdd(t *testing.T, value int, tests []int) {
	for _, in := range tests {
		got := add(in)
		want := value + in
		if got != want {
			t.Errorf("(%d) add(%d) = %d; want %d",
				value,
				in,
				got,
				want)
		}
	}
}
