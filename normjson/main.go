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
