package main

import (
	"strings"
	"testing"
	"time"
)

func TestDaysResult(t *testing.T) {
	test := func(days int, want string) {
		msg := strings.Join([]string{
			"daysResult(%v, %v, %v)",
			"* got *",
			"%v",
			"* want *",
			"%v",
		}, "\n")
		from := parseDate("2012-02-03")
		to := parseDate("2015-02-07")
		if got := daysResult(from, to, days); got != want {
			t.Errorf(msg, from, to, days, got, want)

		}
	}

	test(666, "\nDays:\n666  2013-11-30 (434)")
	test(999, "\nDays:\n999  2014-10-29 (101)")
}

func TestParseDate(t *testing.T) {
	fakeDate := func(y, m, d int) time.Time {
		if y == 0 && m == 0 && d == 0 {
			now := time.Now()
			y = now.Year()
			m = int(now.Month())
			d = now.Day()
		}
		return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	}

	test := func(in string, want time.Time) {
		if got := parseDate(in); got != want {
			t.Errorf("parseDate(\"%v\") = %v; want %v",
				in, formatDate(got), formatDate(want))
		}
	}

	test("abc", fakeDate(0, 0, 0))
	test("2015-02-07", fakeDate(2015, 2, 7))
}
