package embtype

import (
	"testing"
)

func Test(t *testing.T) {
	got := newNamedPoint("one", 1, 2)
	want := "one x:1 y:2"
	if got.String() != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
