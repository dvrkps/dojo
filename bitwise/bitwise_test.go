package main

import "testing"

var result bool

func BenchmarkIsEvenBitwise4(b *testing.B) {
	benchmark(4, isEvenBitwise, b)
}

func BenchmarkIsEvenBitwise5(b *testing.B) {
	benchmark(5, isEvenBitwise, b)
}

func BenchmarkIsEven4(b *testing.B) {
	benchmark(4, isEven, b)
}

func BenchmarkIsEven5(b *testing.B) {
	benchmark(5, isEven, b)
}

func benchmark(i int, f func(i int) bool, b *testing.B) {
	var r bool
	for n := 0; n < b.N; n++ {
		r = f(i)
	}
	result = r
}
