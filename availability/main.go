package availability

import "time"

// Year in nanoseconds.
const yearNs = 3.154e16 * time.Nanosecond

// ByTime holds time based availability data.
//
// formula: availability = uptime / ( uptime + downtime )
type ByTime struct {
	percent  float64
	uptime   time.Duration
	downtime time.Duration
}

// NewByTime creates time based availability.
func NewByTime() *ByTime {
	return &ByTime{downtime: yearNs}
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
