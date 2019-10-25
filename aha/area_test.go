package main

import "testing"

func TestNewArea(t *testing.T) {
	var tests = []struct {
		fail bool
		m2   float64
		ral  float64
		chv  float64
		want string
	}{
		{fail: true, m2: -1, ral: 0, chv: 0},
		{fail: true, m2: 0, ral: -1, chv: 0},
		{fail: true, m2: 0, ral: 0, chv: -1},
		{m2: 1, ral: 0, chv: 0, want: "1 m2 = 0 ral, 0 chv"},
		{m2: 0, ral: 1, chv: 0, want: "5754.542 m2 = 1 ral, 0 chv"},
		{m2: 0, ral: 0, chv: 1, want: "3.596652 m2 = 0 ral, 1 chv"},
		{m2: 12121, ral: 0, chv: 0, want: "12121 m2 = 2 ral, 170 chv"},
	}
	for _, tt := range tests {
		area, err := NewArea(tt.m2, tt.ral, tt.chv)
		if tt.fail {
			if err == nil {
				t.Errorf("NewValue(%v, %v, %v) : nil error",
					tt.m2, tt.ral, tt.chv)
			}
			continue
		}
		got := area.String()
		if got != tt.want || err != nil {
			t.Errorf("NewArea(%v, %v, %v).String() = %q, %v; want %q, <nil>",
				tt.m2, tt.ral, tt.chv,
				got, err,
				tt.want)
		}
	}
}
