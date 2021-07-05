package timecache

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	const want = 1 * time.Second
	tc := New(want)
	got := tc.delayDuration
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
