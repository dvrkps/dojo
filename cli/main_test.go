package main

import (
	"log"
	"testing"
)

func TestRun(t *testing.T) {
	log.SetFlags(0)

	fakeCLI := &CLI{}

	tests := map[string]struct {
		in   int
		want int
	}{

		"ok": {
			in:   1,
			want: 0},

		"in < 1": {
			in:   0,
			want: 1},
	}

	for k, tt := range tests {
		got := run(fakeCLI, tt.in)
		if got != tt.want {
			t.Errorf("%s: run(_, %d) = %d; want %d",
				k, tt.in, got, tt.want)
		}

	}

}
