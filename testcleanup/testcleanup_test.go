package testcleanup

import "testing"

func TestSum(t *testing.T) {
	t.Cleanup(func() {
		t.Log("cleanup")
	})
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{a: 1, b: 2, want: 3},
	}
	for _, tt := range tests {
		got := sum(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("sum(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
