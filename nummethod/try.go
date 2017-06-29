package nummethod

import "reflect"

type FooBar interface {
	foo()
	Bar()
}

type O struct {
	FooBar
}

func Num(o O) int {
	v := reflect.ValueOf(in)
	return v.Field(0).NumMethod()
}
