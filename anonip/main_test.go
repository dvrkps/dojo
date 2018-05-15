package main

import "testing"

func TestAnonymizeIP(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{
			in:   "1.2.3.4",
			want: "1.2.0.0",
		},
		{
			in:   "1:2:3:4:5:6:7:8",
			want: "1:2:3::",
		},
		{
			in:   "::6:7:8",
			want: "::",
		},
		{
			in:   "0.0.3.4",
			want: "0.0.0.0",
		},
		{
			in:   "invalid",
			want: "<nil>",
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
