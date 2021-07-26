package timecache

import (
	"testing"
	"time"
)

func TestDelete(t *testing.T) {
	const key = "first"

	s := newFakeStorage()

	exists := s.Delayed(key, afterEndTime())
	if exists {
		t.Error("exists")
	}

	exists = s.Delayed(key, startTime())
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
		{name: "start", key: "first", now: startTime(), want: true},
		{name: "before", key: "first", now: beforeEndTime(), want: true},
		{name: "end", key: "first", now: endTime(), want: true},
		{name: "notexists", key: "notexists", now: startTime(), want: false},
		{name: "after", key: "first", now: afterEndTime(), want: false},
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
