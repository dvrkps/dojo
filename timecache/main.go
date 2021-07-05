package timecache

import "time"

type TC struct {
	delayDuration time.Duration
}

func New(delayDuration time.Duration) *TC {
	tc := TC{
		delayDuration: delayDuration,
	}

	return &tc
}
