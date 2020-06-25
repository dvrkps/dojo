package isas

import (
	"errors"
	"testing"
)

func TestErrorMessages(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{name: "myError", err: &myError{}, want: "my error"},
		{name: "subError", err: &subError{}, want: "sub error"},
		{name: "myError{subError{sentinelError}}",
			err:  &myError{err: &subError{err: sentinelError}},
			want: "my: sub: sentinel error"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := tt.err.Error()
			if got != tt.want {
				t.Fatalf("got %q; want %q", got, tt.want)
			}
		})
	}
}

func TestIsSentinelError(t *testing.T) {
	err := testFunction()

	if !errors.Is(err, sentinelError) {
		t.Fatalf("not sentinel error")
	}
}

func testFunction() error {
	return &myError{err: &subError{err: sentinelError}}
}

func TestAsSubError(t *testing.T) {
	err := testFunction()

	var serr *subError

	if !errors.As(err, &serr) {
		t.Fatalf("not sub error")
	}
}
