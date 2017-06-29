package nummethod

import "testing"

var testCases = []struct {
	in   interface{}
	want int
}{
	{in: EmbeddedInterface{}, want: 1},
	{in: EmbeddedStruct{}, want: 1},
}

func TestNum(t *testing.T) {
	for _, tc := range testCases {
		got := Num(tc.in)
		if got != tc.want {
			t.Errorf("Num(%T) = %v; want %v", tc.in, got, tc.want)
		}
	}
}
