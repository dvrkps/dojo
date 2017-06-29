package normjson

import (
	"reflect"

	"github.com/dvrkps/dojo/normjson/abc"
)

// T holds data.
type T struct {
	I int64   `json:"i,omitempty"`
	F float64 `json:"f,omitempty"`
	B bool    `json:"b,omitempty"`
	S string  `json:"s,omitempty"`

	Label  interface{} `json:"label,omitempty"`
	Active interface{} `json:"active,omitempty"`
	Number interface{} `json:"number,omitempty"`

	abc.Doer  `json:"-"`
	abc.Empty `json:"-"`
}

func (t *T) normalize() {
	e := reflect.ValueOf(t).Elem()
	var f reflect.Value
	for i, max := 0, e.NumField(); i < max; i++ {
		f = e.Field(i)
		if isInterfaceFieldEmpty(f) {
			f.Set(reflect.Zero(f.Type()))
		}
	}
}

func isInterfaceFieldEmpty(f reflect.Value) bool {
	if f.Kind() != reflect.Interface {
		return false
	}
	if f.NumMethod() != 0 {
		return false
	}
	v := f.Interface()
	all := []interface{}{"", "0", 0, float64(0), int64(0), false, "false"}
	for _, a := range all {
		if v == a {
			return true
		}
	}
	return false
}
