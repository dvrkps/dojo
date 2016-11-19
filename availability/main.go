package availability

import (
	"errors"
	"time"
)

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

// SetPercent sets percent value.
func (a *ByRequests) SetPercent(v float64) error {
	if v <= 0 || v >= 100 {
		return errors.New("invalid percent")
	}
	a.percent = v
	return nil
}

// SetSuccess sets success requests value.
func (a *ByRequests) SetSuccess(v int) error {
	if v <= 0 {
		return errors.New("invalid success requests value")
	}
	a.success = v
	return nil
}

// SetTotal sets total requests value.
func (a *ByRequests) SetTotal(v int) error {
	if v <= 0 {
		return errors.New("invalid total requests value")
	}
	a.total = v
	return nil
}
