package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		in   int
		want int
	}{
		{in: 1, want: 1},
		{in: 2, want: 2},
		{in: 3, want: 3},
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
