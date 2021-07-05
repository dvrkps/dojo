package timecache

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	const delayDuration = 1 * time.Second
	got := New(delayDuration)
	if got == nil {
		t.Error("nil TC")
	}
}
