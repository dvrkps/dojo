package availability

import "time"

// time based formula
// availability = yearUptime / ( yearUptime + yearDowntime )

// aggregate availability formula
// availability = daySuccessfulRequsts / dayTotalRequests

// Year in nanoseconds.
const yearNs = 3.154e16 * time.Nanosecond

// Availability types.
const (
	TimeType = iota
	AggregateType
)
