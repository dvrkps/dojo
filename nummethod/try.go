package nummethod

import "reflect"

type UnExer interface {
	un()
	Ex()
}

type EmbeddedInterface struct {
	UnExer
}

type EmbeddedStruct struct {
	Data
}

type Data struct{}

func (Data) un() {}
func (Data) Ex() {}

func Num(in interface{}) int {
	v := reflect.ValueOf(in)
	return v.Field(0).NumMethod()
}
