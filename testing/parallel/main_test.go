package parallel

import (
	"strconv"
	"testing"
)

const want = 1

func TestParallel(t *testing.T) {
	const max = 2

	for i := 0; i < max; i++ {
		name := strconv.Itoa(i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := Same(want)
			if got != want {
				t.Errorf("got %v; want %v", got, want)
			}
		})
	}
}
