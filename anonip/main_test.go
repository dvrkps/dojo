package main

import "testing"

func TestAnonymizeIP(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{
			in:   "8.8.8.8",
			want: "8.8.0.0",
		},
	}
	for _, tc := range cases {
		got := AnonymizeIP(tc.in)
		if got.String() != tc.want {
			t.Errorf("AnonymizeIP( %s ) = %s; want %s",
				tc.in, got.String(), tc.want)
		}
	}
}
