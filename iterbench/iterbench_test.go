package iterbench

import "testing"

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

func TestIterGoto(t *testing.T) {
	_ = iterGoto(10)
}

var result int

func BenchmarkIterFor(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = iterFor(10)
	}
	result = r
}

func BenchmarkIterGoto(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = iterGoto(10)
	}
	result = r
}
