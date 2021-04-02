package short

import "testing"

func TestSum(t *testing.T) {
	got := sum(1, 2)
	const want = 3
	if got != want {
		t.Errorf("sum(1,2) = %v; want %v", got, want)
	}
}

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test")
	}
	got := sum(1, 2)
	const want = 3
	if got != want {
		t.Errorf("sum(1,2) = %v; want %v", got, want)
	}
}
