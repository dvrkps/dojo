package timecache

import (
	"testing"
	"time"
)

func TestDelete(t *testing.T) {
	const key = "first"

	s := newFakeStorage()

	exists := s.Delayed(key, newFakeDate(1, 2, 14))
	if exists {
		t.Error("exists")
	}

	exists = s.Delayed(key, newFakeDate(1, 2, 3))
	if exists {
		t.Error("not deleted")
	}
}

func TestDelayed(t *testing.T) {
	tests := []struct {
		name string
		key  string
		now  time.Time
		want bool
	}{
		{name: "valid", key: "first", now: newFakeDate(1, 2, 12), want: true},
		{name: "notexists", key: "notexists", now: newFakeDate(1, 2, 13), want: false},
		{name: "delayed", key: "first", now: newFakeDate(1, 2, 14), want: false},
	}

	s := newFakeStorage()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.Delayed(tt.key, tt.now)
			if got != tt.want {
				t.Errorf("key %v got %v", tt.key, got)
			}
		})
	}
}

func newFakeStorage() Storage {
	const tenSecond = 10 * time.Second
	s := New(tenSecond)
	s.Add("first", newFakeDate(1, 2, 3))

	return s
}

func newFakeDate(h, m, s int) time.Time {
	const (
		year  = 1
		month = time.Month(1)
		day   = 1
	)

	return time.Date(year, month, day, h, m, s, 0, time.UTC)
}
