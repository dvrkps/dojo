package iterbench

import (
	"fmt"
	"testing"
)

type testType struct {
	in   int
	want int
}

type testFunc func(int) int

var tests = []testType{
	{10, 55},
	{100, 5050},
	{1000, 500500},
	{10000, 50005000},
	{100000, 5000050000},
}

func testIter(t *testing.T, tests []testType, funcName string, fn testFunc) {
	for _, tt := range tests {
		if got := iterFor(tt.in); got != tt.want {
			t.Errorf("%s(%d) = %d; want %d",
				funcName, tt.in, got, tt.want)
		}
	}
}

func TestIter(t *testing.T) {
	t.Run("iterFor", func(t *testing.T) {
		testIter(t, tests, "iterFor", iterFor)
	})
	t.Run("iterGoto", func(t *testing.T) {
		testIter(t, tests, "iterGoto", iterGoto)
	})
}

var result int

func benchIter(b *testing.B, in int, fn testFunc) {
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = fn(in)
	}
	result = r
}

func BenchmarkIter(b *testing.B) {
	var name string
	for _, tt := range tests {
		name = fmt.Sprintf("iterFor(%d)", tt.in)
		b.Run(name, func(b *testing.B) {
			benchIter(b, tt.in, iterFor)
		})
		name = fmt.Sprintf("iterGoto(%d)", tt.in)
		b.Run(name, func(b *testing.B) {
			benchIter(b, tt.in, iterGoto)
		})
	}
}
