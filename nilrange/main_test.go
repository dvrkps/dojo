package main

import (
	"testing"
)

var testCases = []struct {
	name string
	fn   func([]int) int
	args []int
	want int
}{
	{"one(nil)", one, nil, 0},
	{"one(1,2,3)", one, []int{1, 2, 3}, 6},
	{"two(nil)", two, nil, 0},
	{"two(1,2,3)", two, []int{1, 2, 3}, 6},
	{"three(nil)", three, nil, 0},
	{"three(1,2,3)", three, []int{1, 2, 3}, 6},
}

func TestAll(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.fn(tc.args)
			if got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
