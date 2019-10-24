package main

import (
	"fmt"
	"testing"
)

var result []uint

func pureHelper(b *testing.B, max uint) {
	var r []uint
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = pure(max)
	}
	result = r
}

var pureBenchCases = []struct {
	max uint
}{
	{max: 50},
	{max: 1000},
}

func BenchmarkPure(b *testing.B) {
	for _, bb := range pureBenchCases {
		name := fmt.Sprintf("pure(%v)", bb.max)
		b.Run(name, func(b *testing.B) {
			pureHelper(b, bb.max)
		})
	}
}

func primesHelper(b *testing.B, p primeser) {
	var r []uint
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = p.primes()
	}
	result = r
}

var primesBenchCases = []struct {
	name string
	p    primeser
}{
	{"base(50)", newBase(50)},
	{"base(1000)", newBase(1000)},
}

func BenchmarkPrimes(b *testing.B) {
	for _, bb := range primesBenchCases {
		b.Run(bb.name, func(b *testing.B) {
			primesHelper(b, bb.p)
		})
	}
}
