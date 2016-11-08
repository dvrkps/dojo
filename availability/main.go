package availability

import "time"

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

// ByRequests holds aggregate based availability.
//
// formula: availability = success / total
type ByRequests struct {
	percent float64
	success int
	total   int
}

// NewByRequests creates aggregate based availability.
func NewByRequests() *ByRequests {
	return &ByRequests{}
}
