package timecache

import (
	"sync"
	"time"
)

type TC struct {
	delayDuration time.Duration
	mu            *sync.Mutex
	keyValues     map[string]time.Time
}

func New(delayDuration time.Duration) TC {
	tc := TC{
		delayDuration: delayDuration,
		mu:            &sync.Mutex{},
		keyValues:     make(map[string]time.Time),
	}

	return tc
}
