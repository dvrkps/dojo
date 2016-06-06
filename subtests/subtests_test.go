package main

import "testing"

var sumTests = []struct {
	a, b, want int
}{
	{1, 2, 3},
	{3, 4, 7},
	{5, 6, 11},
}

func TestSum_old(t *testing.T) {
	for _, tt := range sumTests {
		if got := sum(tt.a, tt.b); got != tt.want {
			t.Errorf("sum(%v,%v) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSum_subtests(t *testing.T) {
	for _, tt := range sumTests {
		ok := t.Run("subtest", func(t *testing.T) {
			if got := sum(tt.a, tt.b); got != tt.want {
				t.Errorf("sum(%v,%v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
		if !ok {
			t.FailNow()
			t.Log("subtest fail")
		}
	}
}
