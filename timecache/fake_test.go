package timecache

import (
	"testing"
	"time"
)

func TestTimes(t *testing.T) {
	t.Errorf("start:\n%v", startTime())
	t.Errorf("end:\n%v", endTime())
}

const testDelayDuration = 10 * time.Second

func newFakeStorage() Storage {
	s := New(testDelayDuration)
	s.Add("first", startTime())

	return s
}

func startTime() time.Time {
	return fakeTime(1, 2, 3)
}

func endTime() time.Time {
	s := startTime()
	return s.Add(testDelayDuration)
}

func fakeTime(h, m, s int) time.Time {
	const (
		year  = 1
		month = time.Month(1)
		day   = 1
	)

	return time.Date(year, month, day, h, m, s, 0, time.UTC)
}
