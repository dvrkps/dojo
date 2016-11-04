package availability

import (
	"errors"
	"time"
)

// time based formula
// availability = yearUptime / ( yearUptime + yearDowntime )

// aggregate availability formula
// availability = daySuccessfulRequsts / dayTotalRequests

// Year in nanoseconds.
const yearNs = 3.154e16 * time.Nanosecond

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

// New creates availability.
func New(typ int) (*Availability, error) {
	if typ < TimeType || typ > AggregateType {
		return nil, errors.New("invalid type")
	}
	return &Availability{typ: typ}, nil
}
