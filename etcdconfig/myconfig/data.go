package myconfig

import (
	"sync"
	"sync/atomic"
)

type dataMap map[string]string

type data struct {
	mu sync.Mutex
	av atomic.Value
}

func newData() *data {
	d := &data{}
	d.av.Store(dataMap{})
	return d
}

func (d *data) all() dataMap {
	return d.av.Load().(dataMap)
}

func (d *data) get(key string) (string, bool) {
	m := d.av.Load().(dataMap)
	v, ok := m[key]
	return v, ok
}

func (d *data) update(dm dataMap) {
	d.mu.Lock()
	defer d.mu.Unlock()

	fresh := make(dataMap)

	// clone old values
	old := d.av.Load().(dataMap)
	for k, v := range old {
		fresh[k] = v
	}

	// update values
	for k, v := range dm {
		fresh[k] = v
	}

	d.av.Store(fresh)
}
