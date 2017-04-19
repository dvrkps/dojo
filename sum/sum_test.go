package sum

import "testing"

var testCases = []struct {
	name string
	fn   func(int, ...int) bool
	sum  int
	in   []int
	want bool
}{
	{name: "basic(8,1,2,3,9)",
		fn:   basic,
		sum:  8,
		in:   []int{1, 2, 3, 9},
		want: false},
	{name: "basic(8,1,2,4,4)",
		fn:   basic,
		sum:  8,
		in:   []int{1, 2, 4, 4},
		want: true},
	{name: "better(8,1,2,3,9)",
		fn:   better,
		sum:  8,
		in:   []int{1, 2, 3, 9},
		want: false},
	{name: "better(8,1,2,4,4)",
		fn:   better,
		sum:  8,
		in:   []int{1, 2, 4, 4},
		want: true},
	{name: "linear(8,1,2,3,9)",
		fn:   linear,
		sum:  8,
		in:   []int{1, 2, 3, 9},
		want: false},
	{name: "linear(8,1,2,4,4)",
		fn:   linear,
		sum:  8,
		in:   []int{1, 2, 4, 4},
		want: true},
}

func TestAll(t *testing.T) {
	for _, tc := range testCases {
		got := tc.fn(tc.sum, tc.in...)
		if got != tc.want {
			t.Errorf("%s = %v; want %v", tc.name, got, tc.want)
		}
	}
}
