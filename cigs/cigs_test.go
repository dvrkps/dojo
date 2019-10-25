package main

import (
	"testing"
	"time"
)

func formatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

func TestDateSince(t *testing.T) {
	since := parseDate("2012-02-03")
	var tests = []struct {
		days int
		want time.Time
	}{
		{777, parseDate("2014-03-21")},
		{888, parseDate("2014-07-10")},
		{999, parseDate("2014-10-29")},
	}
	for _, tt := range tests {
		if got := DateSince(since, tt.days); got != tt.want {
			t.Errorf("DateSince(%v, %v) = %v; want %v",
				formatDate(since),
				tt.days,
				formatDate(got),
				formatDate(tt.want),
			)
		}
	}
}

func TestDaysBetween(t *testing.T) {
	since := parseDate("2012-02-03")
	var tests = []struct {
		now  time.Time
		want int
	}{
		{parseDate("2014-03-21"), 777},
		{parseDate("2014-07-10"), 888},
		{parseDate("2014-10-29"), 999},
	}
	for _, tt := range tests {
		if got := DaysBetween(since, tt.now); got != tt.want {
			t.Errorf("DaysBetween(%v, %v) = %v; want %v",
				formatDate(since),
				formatDate(tt.now), got, tt.want)
		}
	}
}
