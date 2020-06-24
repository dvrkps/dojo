package tlog

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "valid", a: 1, b: 2, want: 3},
		{name: "invalid", a: 1, b: 2, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sum(tt.a, tt.b)
			t.Logf("sum(%v, %v) = %v", tt.a, tt.b, got)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
