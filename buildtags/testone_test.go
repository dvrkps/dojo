// +build testone

package main

import "testing"

func TestOneAdd(t *testing.T) {
	tests := []struct {
		in   int
		want int
	}{
		{in: 11, want: 11},
		{in: 12, want: 12},
		{in: 13, want: 13},
	}
	for _, tt := range tests {
		got := add(tt.in)
		if got != tt.want {
			t.Errorf("(%d) add(%d) = %d; want %d",
				value,
				tt.in,
				got,
				tt.want)
		}
	}

}
