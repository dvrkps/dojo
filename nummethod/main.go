package main

import (
	"fmt"
	"reflect"
)

type UnExer interface {
	un()
	Ex()
}

type embeddedInterface struct {
	UnExer
}

type embeddedStruct struct {
	data
}

type data struct{}

func (data) un() {}

func (data) Ex() {}

func num(in interface{}) int {
	v := reflect.ValueOf(in)
	return v.Field(0).NumMethod()
}

func main() {
	var try = []struct {
		in   interface{}
		want int
	}{
		{in: embeddedInterface{}, want: 1},
		{in: embeddedStruct{}, want: 1},
	}

	for _, tc := range try {
		got := num(tc.in)
		if got != tc.want {
			fmt.Printf("num(%T) = %v; want %v\n", tc.in, got, tc.want)
		}
	}
}
