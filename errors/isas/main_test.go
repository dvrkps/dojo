package isas

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsSentinelError(t *testing.T) {
	fn := func() error {
		return &myError{err: &myError{err: sentinelError}}
	}

	err := fn()

	if !errors.Is(err, sentinelError) {
		t.Fatalf("not sentinel error")
	}

	got := fmt.Sprintf("%v", err)

	const want = "my: my: sentinel error"

	if got != want {
		t.Fatalf("got %q; want %q", got, want)
	}
}
