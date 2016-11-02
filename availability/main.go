package availability

// time based formula
// availability = yearUptime / ( yearUptime + yearDowntime )

// aggregate availability formula
// availability = daySuccessfulRequsts / dayTotalRequests

const version = "0.1.0"

// Availability types.
const (
	TimeType = iota
	AggregateType
)
