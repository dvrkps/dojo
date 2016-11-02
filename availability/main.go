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

// Availability is data holder.
type Availability struct {
	typ int

	percent float64

	yearUptime   time.Duration
	yearDowntime time.Duration

	daySuccessfulRequests int
	dayTotalRequests      int
}
