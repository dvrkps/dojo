package timecache

import (
	"time"
)

const testDelayDuration = 1 * time.Hour

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

func afterEndTime() time.Time {
	e := endTime()
	return e.Add(1 * time.Second)
}

func beforeEndTime() time.Time {
	e := endTime()
	return e.Add(-1 * time.Second)
}

func fakeTime(h, m, s int) time.Time {
	const (
		year  = 1
		month = time.Month(1)
		day   = 1
	)

	return time.Date(year, month, day, h, m, s, 0, time.UTC)
}
