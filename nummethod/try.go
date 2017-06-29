package main

import "reflect"

type UnExer interface {
	un()
	Ex()
}

type embeddedInterface struct {
	UnExer
}

type embeddedStruct struct {
	Data
}

type data struct{}

func (data) un() {}
func (data) Ex() {}

func ValueNum(in interface{}) int {
	v := reflect.ValueOf(in)
	return v.Field(0).NumMethod()
}
