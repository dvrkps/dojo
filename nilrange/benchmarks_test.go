package main

import (
	"fmt"
	"testing"
)

var result int

func benchmarkHelper(b *testing.B, fn func([]int) int, args []int) {
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = fn(args)
	}
	result = r
}

var benchCases = []struct {
	name string
	fn   func([]int) int
	args []int
}{
	{"one(nil)", one, nil},
	{"one(1,2,3)", one, []int{1, 2, 3}},
}

func BenchmarkAll(b *testing.B) {
	for _, bc := range benchCases {
		name := fmt.Sprintf("%s(%v)", bc.name, bc.args)
		b.Run(name, func(b *testing.B) {
			benchmarkHelper(b, bc.fn, bc.args)
		})
	}
}
