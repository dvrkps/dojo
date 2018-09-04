package special

import (
	"testing"
)

var result bool

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
