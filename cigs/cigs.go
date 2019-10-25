package main

import (
	"math"
	"time"
)

// DateSince returns date for days since date.
func DateSince(from time.Time, days int) time.Time {
	// convert absolute days to hours
	h := int(math.Abs(float64(days)) * 24)
	// calc duration in hours
	dur := time.Duration(h) * time.Hour
	return from.Add(dur)
}

// DaysBetween returns days between dates.
func DaysBetween(from, to time.Time) int {
	// convert diff hours to days
	d := to.Sub(from).Hours() / 24
	return int(math.Abs(d))
}
