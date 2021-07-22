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

func TestAdd(t *testing.T) {
	const second = 1 * time.Second
	tc := New(second)
	now := testDate(1, 2, 3)
	const key = "key"
	tc.Add(key, now)

	got := tc.Delayed(key, testDate(1, 2, 3))
	want := true
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}

}

func testDate(h, m, s int) time.Time {
	const (
		year  = 1
		month = time.Month(1)
		day   = 1
	)

	return time.Date(year, month, day, h, m, s, 0, time.UTC)
}
