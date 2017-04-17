package main

import (
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
	{"two(nil)", two, nil},
	{"three(nil)", three, nil},
	// not nil
	{"one(1,2,3)", one, []int{1, 2, 3}},
	{"two(1,2,3)", two, []int{1, 2, 3}},
	{"three(1,2,3)", three, []int{1, 2, 3}},
}

func BenchmarkAll(b *testing.B) {
	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			benchmarkHelper(b, bc.fn, bc.args)
		})
	}
}
