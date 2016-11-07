package availability

import "time"

// aggregate availability formula
// availability = daySuccessfulRequsts / dayTotalRequests

// Availability types.
const (
	TimeType = iota + 1
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

// Year in nanoseconds.
const yearNs = 3.154e16 * time.Nanosecond

// TimeBased holds time based availability data.
//
// formula: availability = yearUptime / ( yearUptime + yearDowntime )
type TimeBased struct {
	percent      float64
	yearUptime   time.Duration
	yearDowntime time.Duration
}

// NewTimeBased creates time based availability.
func NewTimeBased() *TimeBased {
	return &TimeBased{yearDowntime: yearNs}
}

// ByAggregate holds aggregate based availability.
//
// formula: availability = daySuccessfulRequsts / dayTotalRequests
type ByAggregate struct {
	percent            float64
	successfulRequests int
	totalRequests      int
}
