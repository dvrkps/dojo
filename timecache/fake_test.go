package timecache

import "time"

func newFakeStorage() Storage {
	const tenSecond = 10 * time.Second
	s := New(tenSecond)
	s.Add("first", newBaseTime())

	return s
}

func newBaseTime() time.Time {
	return newFakeDate(1, 2, 3)
}

func newFakeDate(h, m, s int) time.Time {
	const (
		year  = 1
		month = time.Month(1)
		day   = 1
	)

	return time.Date(year, month, day, h, m, s, 0, time.UTC)
}