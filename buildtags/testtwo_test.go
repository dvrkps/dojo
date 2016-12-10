// +build testtwo

package main

import "testing"

func TestTwoAdd(t *testing.T) {
	tests := []struct {
		in   int
		want int
	}{
		{in: 21, want: 21},
		{in: 22, want: 22},
		{in: 23, want: 23},
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
