package iterbench

import "testing"

func TestIterFor(t *testing.T) {
	_ = iterFor(10)
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
