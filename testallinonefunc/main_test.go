package main

import "testing"

func TestInc(t *testing.T) {
	tests := map[string]struct {
		in    int
		want  int
		isErr bool
	}{
		"ok":     {in: 1, want: 2, isErr: false},
		"== max": {in: incMax, want: incMax + 1, isErr: false},
		"< 0":    {in: -1, want: 0, isErr: true},
		"> max":  {in: incMax + 1, want: 0, isErr: true},
	}

	const msg = "%s: inc(%d) = %d, %v; want %d, %v"

	for k, tt := range tests {

		got, err := inc(tt.in)

		if tt.isErr {

			if got != tt.want || err == nil {
				t.Errorf(msg, k, tt.in, got, err, tt.want, "<error>")
			}

			continue
		}

		if got != tt.want || err != nil {

			t.Errorf(msg, k, tt.in, got, err, tt.want, "<nil>")

		}
	}

}
