package short

import "testing"

func TestSum(t *testing.T) {
	got := sum(1, 2)
	const want = 3
	if got != want {
		t.Errorf("sum(1,2) = %v; want %v", got, want)
	}
}
