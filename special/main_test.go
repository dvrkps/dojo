package special

import (
	"testing"
)

var result bool

func TestSpecial(t *testing.T) {
	for i := uint32(0); i < 1000000000; i++ {
		want := naive(i)

		if got := lookup(i); got != want {
			t.Errorf("lookup(%v) = %v; want %v", i, got, want)
		}
		if got := leftShift(i); got != want {
			t.Errorf("leftShift(%v) = %v; want %v", i, got, want)
		}
	}
}

func BenchmarkNaive(b *testing.B) {
	var r bool
	const x = 11
	for n := 0; n < b.N; n++ {
		r = naive(11)
	}
	result = r
}

func BenchmarkLookup(b *testing.B) {
	var r bool
	const x = 11
	for n := 0; n < b.N; n++ {
		r = lookup(11)
	}
	result = r
}

func BenchmarkLeftShift(b *testing.B) {
	var r bool
	const x = 11
	for n := 0; n < b.N; n++ {
		r = leftShift(11)
	}
	result = r
}
