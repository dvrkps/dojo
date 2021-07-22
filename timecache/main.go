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

func (tc *TC) Add(key string, now time.Time) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.keyValues[key] = now.Add(tc.delayDuration)
}

func (tc *TC) Delayed(key string, now time.Time) bool {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	got, exists := tc.keyValues[key]
	if !exists {
		return false
	}

	expired := now.After(got)
	if expired {
		delete(tc.keyValues, key)
		return false
	}

	return true
}
