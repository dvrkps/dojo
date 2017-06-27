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

	abc.Doer
}

func (t *T) normalize() {
	e := reflect.ValueOf(t).Elem()
	var f reflect.Value
	for i, max := 0, e.NumField(); i < max; i++ {
		f = e.Field(i)
		if f.Kind() != reflect.Interface || f.NumMethod() != 0 {
			continue
		}
		//if f.Type().String() != "interface {}" {
		//	continue
		//}

		v := f.Interface()
		if v == "" ||
			v == "undefined" ||
			v == "null" ||
			v == "NaN" ||
			v == "0" ||
			v == float64(0) ||
			v == int64(0) ||
			v == 0 ||
			v == false ||
			v == "false" {
			f.Set(reflect.Zero(f.Type()))
		}
		//fmt.Printf("%q %q = %v\n",
		//	et.Field(i).Name, f.Type(), f.Interface())
	}
}

func old(t T) T {
	fields := []interface{}{
		t.Label,
		t.Active,
		t.Number,
	}

	for _, v := range fields {
		switch v.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			if v == 0 {
				v = nil
			}
		case float32, float64:
			if v == 0 {
				v = nil
			}
		case string:
			if v == "" {
				v = nil
			}
		}

		//if *v == 0 || *v == "0" || *v == "false" || *v == false || *v == "" {
		//	*v = nil
		//}
	}
	return t
}
