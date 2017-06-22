package testhelper

import "testing"

var testCases = []struct {
	in   []int
	want int
}{
	{in: []int{1, 2}, want: 1},
}

func testRange(t *testing.T, fn func(t *testing.T, got, want int)) {
	for _, tc := range testCases {
		got := Sum(tc.in...)
		fn(t, got, tc.want)
	}
}

func one(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("one: got %v; want %v", got, want)
	}
}

func two(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("two: got %v; want %v", got, want)
	}
}

func TestOne(t *testing.T) {
	testRange(t, one)
}
