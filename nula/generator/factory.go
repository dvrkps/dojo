package main

import (
	"fmt"
	"time"
)

// Factory is JobID generator.
type Factory struct {
	// 0 < x < factoryLimit
	count int
}

// NewFactory creates Factory.
func NewFactory() Factory {
	return Factory{}
}

// NewJobID creates JobID.
func (f *Factory) NewJobID(now time.Time, sid int) string {

	f.next()

	// TODO(dvrkps): better formating
	id := fmt.Sprintf("%d%d%d",
		epochMs(now),
		f.count,
		sid)

	return id
}

const factoryLimit = 99

func (f *Factory) next() int {
	f.count++
	if f.count > factoryLimit {
		f.count = 1
	}
	return f.count
}

func epochMs(t time.Time) int64 {
	return t.UnixNano() / 1000000
}
