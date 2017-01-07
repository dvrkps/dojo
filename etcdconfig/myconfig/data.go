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

func newStorage() *data {
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

	fresh := make(Map)

	// clone old values
	old := d.av.Load().(Map)
	for k, v := range old {
		fresh[k] = v
	}

	// update values
	for k, v := range dm {
		fresh[k] = v
	}

	d.av.Store(fresh)
}

/*
type Map map[string]string
    var m Value
    m.Store(make(Map))
    var mu sync.Mutex // used only by writers
    // read function can be used to read the data without further synchronization
    read := func(key string) (val string) {
            m1 := m.Load().(Map)
            return m1[key]
    }
    // insert function can be used to update the data without further synchronization
    insert := func(key, val string) {
            mu.Lock() // synchronize with other potential writers
            defer mu.Unlock()
            m1 := m.Load().(Map) // load current value of the data structure
            m2 := make(Map)      // create a new value
            for k, v := range m1 {
                    m2[k] = v // copy all data from the current object to the new one
            }
            m2[key] = val // do the update that we need
            m.Store(m2)   // atomically replace the current object with the new one
            // At this point all new readers start working with the new version.
            // The old version will be garbage collected once the existing readers
            // (if any) are done with it.
    }*/
