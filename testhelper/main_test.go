package testhelper

import "testing"

var testCases = []struct {
	in   []int
	want int
}{
	{in: []int{1, 2}, want: 3},
}

func testRange(t *testing.T, fn func(t *testing.T, got, want int)) {
	for _, tc := range testCases {
		got := Sum(tc.in...)
		fn(t, got, tc.want)
	}
}
