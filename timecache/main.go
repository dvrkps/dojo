package timecache

import (
	"sync"
	"time"
)

type Storage struct {
	delayDuration time.Duration
	mu            *sync.Mutex
	keyValues     map[string]time.Time
}

func New(delayDuration time.Duration) Storage {
	s := Storage{
		delayDuration: delayDuration,
		mu:            &sync.Mutex{},
		keyValues:     make(map[string]time.Time),
	}

	return s
}

func (s *Storage) Add(key string, now time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.keyValues[key] = now.Add(s.delayDuration)
}

func (s *Storage) Delayed(key string, now time.Time) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	got, exists := s.keyValues[key]
	if !exists {
		return false
	}

	expired := now.After(got)
	if expired {
		delete(s.keyValues, key)
		return false
	}

	return true
}
