package myconfig

import (
	"sync"
	"sync/atomic"
)

type dataMap map[string]string

type data struct {
	av atomic.Value

	mu sync.Mutex
	dm dataMap
}

func newData() *data {
	d := &data{}
	d.av.Store(d.dm)

	return d
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
